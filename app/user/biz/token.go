package biz

import (
	"eshop/app/user/internal/innerr"
	"eshop/app/user/internal/model"
)

func (b *BusinessHandler) CreateToken(token *model.Token) error {
	_, err := b.data.Db.Insert(token)
	return err
}

func (b *BusinessHandler) GetToken(accessToken string) (token *model.Token, err error) {
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
	_, err = b.data.Db.Where("access_token=?", accessToken).Delete(&model.Token{})
	return
}
