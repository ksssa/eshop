package model

type Token struct {
	Id                 int64  `xorm:"pk"`
	UserId             int64  `xorm:"index default 0 comment('用户id')" json:"userId"`
	AccessToken        string `xorm:"index default '' comment('用户token')" json:"token"`
	AccessTokenExpire  int64  `xorm:"default 0 comment('token有效时间')" json:"accessTokenExpire"`
	RefreshToken       string `xorm:"index default '' comment('刷新token')" json:"refreshToken"`
	RefreshTokenExpire string `xorm:"default 0 comment('刷新token有效时间'')" json:"refreshTokenExpire"`
	CreateTime         int64  `xorm:"created" json:"createTime"`
	UpdateTime         int64  `xorm:"updated" json:"updateTime"`
	DeleteTime         int64  `xorm:"deleted"`
}
