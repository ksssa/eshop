package biz

import (
	user "eshop/api/user/v1"
	"eshop/app/user/internal/innerr"
	"eshop/app/user/internal/model"
)

func (b *BusinessHandler) CreateUser(user *model.User) error {
	_, err := b.data.Db.Insert(user)
	return err
}

func (b *BusinessHandler) UpdateUser(user *model.User) error {
	_, err := b.data.Db.Update(user)
	return err

}

func (b *BusinessHandler) GetUser(id int64) (user *model.User, err error) {
	user = new(model.User)
	exist, err := b.data.Db.ID(id).Get(user)
	if err != nil {
		return
	}
	if exist == false {
		err = innerr.ErrDataNotExist
		return
	}
	return
}

func (b *BusinessHandler) ListUser(filter *user.ListFilter) (users []*model.User, total int64, err error) {
	session := b.data.Db.NewSession()
	if filter.Email != "" {
		session = session.And("email=?", filter.Email)
	}
	if filter.Nick != "" {
		session = session.And("nick=?", filter.Nick)
	}
	start, end := b.data.Page(filter.Page)
	users = make([]*model.User, 0)
	session = session.Limit(end, start).Desc("create_time")
	total, err = session.FindAndCount(&users)
	if err != nil {
		return
	}
	return
}

func (b *BusinessHandler) DeleteUser(id int64) error {
	_, err := b.data.Db.ID(id).Delete(&model.User{})
	return err
}
