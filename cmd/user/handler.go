package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"tiktok/cmd/user/service"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/userpart"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *userpart.UserRegisterRequest) (resp *userpart.UserRegisterResponse, err error) {
	resp = new(userpart.UserRegisterResponse)
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode = 1
		err = errno.UserRegisterFailedErr
		return
	}
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *userpart.UserLoginRequest) (resp *userpart.UserLoginResponse, err error) {
	logrus.Info("The user login")
	resp = new(userpart.UserLoginResponse)
	resp.StatusCode = 0
	success := service.NewLoginUserService(ctx).LoginUser(req)
	if !success {
		resp.StatusCode = 1
		return resp, errno.AuthorizationFailedErr
	}
	return
}
