package service

import (
	"context"
	user "eshop/api/user/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) *errors.Error {

	return nil
}

func (s *UserService) Update(ctx context.Context, req *user.UpdateRequest) *errors.Error {
	return nil
}

func (s *UserService) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, *errors.Error) {
	return nil, nil
}

func (s *UserService) List(ctx context.Context, req *user.ListRequest) (*user.ListResponse, *errors.Error) {
	return nil, nil
}

func (s *UserService) Delete(ctx context.Context, req *user.DeleteRequest) *errors.Error {
	return nil
}
