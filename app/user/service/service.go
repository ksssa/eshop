package service

import (
	user "eshop/api/user/v1"
	"eshop/app/user/biz"
	"eshop/app/user/config"
	"sync"
)

type UserService struct {
	user.UnimplementedUserServer
	bizHandler  *biz.BusinessHandler
	config      *config.Config
	initialized bool
	sync.Mutex
}

func New(config *config.Config) *UserService {
	return &UserService{
		config:      config,
		initialized: false,
	}
}

func (s *UserService) Initialize() error {
	s.Lock()
	defer s.Unlock()
	if s.initialized {
		return nil
	}
	b := biz.New(s.config)
	s.bizHandler = b
	s.bizHandler.Initialized()
	return nil
}
