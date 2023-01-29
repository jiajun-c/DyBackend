package main

import (
	"context"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/cmd/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *userpart.UserRegisterRequest) (resp *userpart.UserRegisterResponse, err error) {
	resp = new(userpart.UserRegisterResponse)
	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.StatusCode = 1
	}
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *userpart.UserLoginRequest) (resp *userpart.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}
