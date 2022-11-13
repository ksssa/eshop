package service

import (
	"context"
	user "eshop/api/user/v1"
	"eshop/app/user/internal/model"
	"eshop/app/user/internal/util"
	"github.com/go-kratos/kratos/v2/errors"
)

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) *errors.Error {

	if len(req.Mobile) == 0 || len(req.Email) == 0 || len(req.Passwd) == 0 {
		return user.ErrorParamNotEnough("mobile,email or passwd is empty")
	}
	if !util.CheckMobileValid(req.Mobile) {
		return user.ErrorParamInvalid("mobile is invalid")
	}
	if !util.CheckPasswdValid(req.Passwd) {
		return user.ErrorParamInvalid("passwd is invalid")
	}
	if !util.CheckEmailValid(req.Email) {
		return user.ErrorParamInvalid("email is invalid")
	}
	passwd, err := util.CryptoPasswd(req.Passwd)
	if err != nil {
		return user.ErrorUnknownError(err.Error())
	}
	userInfo := &model.User{
		Id:     util.GenerateID(),
		Email:  req.Email,
		Nick:   req.Nick,
		Passwd: passwd,
		Mobile: req.Mobile,
	}
	err = s.bizHandler.CreateUser(userInfo)
	if err != nil {
		return user.ErrorUnknownError(err.Error())
	}
	return nil
}

func (s *UserService) Update(ctx context.Context, req *user.UpdateRequest) *errors.Error {
	token, err := util.GetTokenFromCtx(ctx)
	if err != nil {
		return user.ErrorUnknownError(err.Error())
	}
	userId, _, err := s.bizHandler.CheckToken(token)
	if err != nil {
		return user.ErrorTokenInvalid(err.Error())
	}
	if req.User == nil {
		return user.ErrorParamNotEnough("user or id is empty")
	}
	userInfo := model.Pb2User(req.User)
	userInfo.Id = userId
	err = s.bizHandler.UpdateUser(userInfo)
	if err != nil {
		return user.ErrorUnknownError(err.Error())
	}
	return nil
}

func (s *UserService) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, *errors.Error) {
	resp := new(user.GetResponse)
	token, err := util.GetTokenFromCtx(ctx)
	if err != nil {
		return resp, user.ErrorUnknownError(err.Error())
	}
	userId, _, err := s.bizHandler.CheckToken(token)
	if err != nil {
		return resp, user.ErrorTokenInvalid(err.Error())
	}
	userInfo, err := s.bizHandler.GetUser(userId)
	if err != nil {
		return resp, user.ErrorUnknownError(err.Error())
	}
	resp = &user.GetResponse{User: &user.UserInfo{
		Id:    userInfo.Id,
		Nick:  userInfo.Nick,
		Email: userInfo.Email,
		//Name:       userInfo.Nick,
		PublicKey:  userInfo.PublicKey,
		PrivateKey: userInfo.PrivateKey,
		//Status:     userInfo.Status,
		CreateTime: userInfo.CreateTime,
		UpdateTime: userInfo.UpdateTime,
	}}
	return resp, nil
}

func (s *UserService) List(ctx context.Context, req *user.ListRequest) (*user.ListResponse, *errors.Error) {
	//admin才需要
	resp := new(user.ListResponse)
	token, err := util.GetTokenFromCtx(ctx)
	if err != nil {
		return resp, user.ErrorUnknownError(err.Error())
	}
	_, _, err = s.bizHandler.CheckToken(token)
	if err != nil {
		return resp, user.ErrorTokenInvalid(err.Error())
	}
	list, total, err := s.bizHandler.ListUser(req.Filter)
	if err != nil {
		return resp, user.ErrorUnknownError(err.Error())
	}
	users := model.User2PB(list)
	resp.Total = total
	resp.Users = users
	return resp, nil
}

func (s *UserService) Delete(ctx context.Context, req *user.DeleteRequest) *errors.Error {
	//admin才需要
	token, err := util.GetTokenFromCtx(ctx)
	if err != nil {
		return user.ErrorUnknownError(err.Error())
	}
	_, isAdmin, err := s.bizHandler.CheckToken(token)
	if err != nil {
		return user.ErrorTokenInvalid(err.Error())
	}
	if !isAdmin {

	}
	if req.Id == 0 {
		return user.ErrorParamNotEnough("id is empty")
	}
	err = s.bizHandler.DeleteUser(req.Id)
	if err != nil {
		return user.ErrorUnknownError(err.Error())
	}
	return nil
}
