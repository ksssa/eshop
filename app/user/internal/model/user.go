package model

import user "eshop/api/user/v1"

type User struct {
	Id         int64  `xorm:"pk" json:"id"`
	Email      string `xorm:"unique index default '' comment('邮箱')" json:"email"`
	Nick       string `xorm:"default '' comment('昵称')" json:"nick"`
	Passwd     string `xorm:"default '' comment('密码')" json:"passwd"`
	Status     int    `xorm:"default 1 comment('状态')" json:"status"`
	Mobile     string `xorm:"default '' comment('手机号')" json:"mobile"`
	PublicKey  string `xorm:"default '' comment('用户公钥')" json:"publicKey"`
	PrivateKey string `xorm:"default '' comment('用户私钥')" json:"privateKey"`
	CreateTime int64  `xorm:"created" json:"createTime"`
	UpdateTime int64  `xorm:"updated" json:"updateTime"`
}

func User2PB(users []*User) []*user.UserInfo {
	pbUsers := make([]*user.UserInfo, len(users))
	for k, v := range pbUsers {
		u := &user.UserInfo{
			Id:         v.Id,
			Nick:       v.Nick,
			Email:      v.Email,
			Name:       v.Name,
			PublicKey:  v.PublicKey,
			PrivateKey: v.PrivateKey,
			Status:     v.Status,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
		}
		pbUsers[k] = u
	}
	return pbUsers
}

func Pb2User(user *user.UserInfo) *User {
	return &User{
		Id:    user.Id,
		Email: user.Email,
		Nick:  user.Nick,
		//Passwd:     user.Id,
		//Status:     user.Status,
		//Mobile:     user.,
		PublicKey:  user.PublicKey,
		PrivateKey: user.PrivateKey,
	}
}
