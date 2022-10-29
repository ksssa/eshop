package service

import (
	user "eshop/api/user/v1"
	"eshop/app/user/biz"
	"sync"
)

type UserService struct {
	user.UnimplementedUserServer
	bizHandler *biz.BusinessHandler
	sync.Mutex
}

func New() *UserService {
	return nil
}

func (s *UserService) Initialize() {

}
