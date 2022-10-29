package service

import (
	"context"
	user "eshop/api/user/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

func (s *UserService) RefreshToken(ctx context.Context, req *user.RefreshTokenRequest) (*user.RefreshTokenResponse, *errors.Error) {
	return nil, nil
}

func (s *UserService) LoginOut(ctx context.Context, req *user.LoginOutRequest) (*user.LoginOutResponse, *errors.Error) {
	return nil, nil
}

func (s *UserService) Login(ctx context.Context, req *user.LoginResponse) (*user.LoginResponse, *errors.Error) {
	return nil, nil
}
