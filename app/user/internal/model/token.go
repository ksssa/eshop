package model

import "strconv"

type Token struct {
	Id                 int64  `xorm:"pk"`
	IsAdmin            bool   `xorm:"default false comment('是否为管理员')" json:"isAdmin"`
	UserId             int64  `xorm:"index default 0 comment('用户id')" json:"userId"`
	AccessToken        string `xorm:"index default '' comment('用户token')" json:"token"`
	AccessTokenExpire  int64  `xorm:"default 0 comment('token有效时间')" json:"accessTokenExpire"`
	RefreshToken       string `xorm:"index default '' comment('刷新token')" json:"refreshToken"`
	RefreshTokenExpire int64  `xorm:"default 0 comment('刷新token有效时间'')" json:"refreshTokenExpire"`
	CreateTime         int64  `xorm:"created" json:"createTime"`
	UpdateTime         int64  `xorm:"updated" json:"updateTime"`
}

func Token2Cache(token *Token) map[string]interface{} {
	values := make(map[string]interface{})
	values["id"] = token.Id
	values["accessToken"] = token.AccessToken
	values["accessTokenExpire"] = token.AccessTokenExpire
	values["refreshToken"] = token.RefreshToken
	values["userId"] = token.UserId
	values["refreshTokenExpire"] = token.RefreshTokenExpire
	values["isAdmin"] = token.IsAdmin
	return values
}
func Cache2Token(values map[string]string) (*Token, error) {
	token := new(Token)
	id, err := strconv.ParseInt(values["id"], 10, 64)
	if err != nil {
		return token, err
	}
	token.Id = id
	token.AccessToken = values["accessToken"]
	num, err := strconv.ParseInt(values["accessTokenExpire"], 10, 32)
	if err != nil {
		return token, err
	}
	token.AccessTokenExpire = num
	token.RefreshToken = values["refreshToken"]
	num, err = strconv.ParseInt(values["refreshTokenExpire"], 10, 32)
	if err != nil {
		return token, err
	}
	token.RefreshTokenExpire = num
	num, err = strconv.ParseInt(values["userId"], 10, 64)
	if err != nil {
		return token, err
	}
	token.UserId = num
	isTrue, err := strconv.ParseBool(values["isAdmin"])
	if err != nil {
		return token, err
	}
	token.IsAdmin = isTrue
	return token, nil
}
