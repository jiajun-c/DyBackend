package main

import (
	"context"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/cmd/user/service"
	"tiktok/internal/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *userpart.UserLoginRequest) (resp *userpart.UserLoginResponse, err error) {
	resp = &userpart.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success to login",
		UserId:     0,
		Token:      "",
	}
	success := service.NewLoginUserService(ctx).LoginUser(req)
	if !success {
		resp.StatusCode = 1
	}
	return resp, nil
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *userpart.UserRegisterRequest) (resp *userpart.UserRegisterResponse, err error) {
	resp = &userpart.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success to register",
		UserId:     0,
		Token:      "",
	}
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "register failed"
		err = errno.UserRegisterFailedErr
		return
	}
	return
}
