package biz

import (
	"eshop/app/user/internal/innerr"
	"eshop/app/user/internal/model"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-xorm/xorm"
	"time"
)

func (b *BusinessHandler) CreateToken(token *model.Token) error {
	_, err := b.data.Db.Insert(token)
	key := fmt.Sprintf("%s%s", model.CacheTokenPrefix, token.AccessToken)
	b.data.Cache.Del(key)
	b.cacheToken(token)
	return err
}

func (b *BusinessHandler) getToken(accessToken string) (token *model.Token, err error) {
	token = new(model.Token)
	exist, err := b.data.Db.Where("access_token=?", accessToken).Get(token)
	if err != nil {
		return
	}
	if exist == true {
		err = innerr.ErrDataNotExist
		return
	}
	return
}

func (b *BusinessHandler) DeleteToken(accessToken string) (err error) {
	session := b.data.Db.NewSession()
	err = b.data.Transaction(session, func(session *xorm.Session) error {
		_, err = b.data.Db.Where("access_token=?", accessToken).Delete(&model.Token{})
		if err != nil {
			return err
		}
		key := fmt.Sprintf("%s%s", model.CacheTokenPrefix, accessToken)
		_, err = b.data.Cache.Del(key).Result()
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (b *BusinessHandler) CheckToken(token string) (int64, bool, error) {
	//返回是否管理员
	tokenInfo, err := b.getTokenFromCache(token)
	if err != nil && err != innerr.ErrCacheNotExist {
		return 0, false, err
	}
	if err == innerr.ErrCacheNotExist {
		tokenInfo, err = b.getToken(token)
		if err != nil {
			log.Error(err)
			return 0, false, err
		}
	}
	if tokenInfo.AccessTokenExpire < time.Now().Unix() {
		return 0, false, innerr.ErrTokenExpire
	}
	return tokenInfo.UserId, tokenInfo.IsAdmin, nil
}

func (b *BusinessHandler) cacheToken(token *model.Token) error {
	key := fmt.Sprintf("%s%s", model.CacheTokenPrefix, token.AccessToken)
	_, err := b.data.Cache.HMSet(key, model.Token2Cache(token)).Result()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (b *BusinessHandler) getTokenFromCache(token string) (*model.Token, error) {

	key := fmt.Sprintf("%s%s", model.CacheTokenPrefix, token)
	res, err := b.data.Cache.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, innerr.ErrCacheNotExist
	}
	tokenInfo, err := model.Cache2Token(res)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
