package main

import (
	"context"
	user_part "tiktok/cmd/user/kitex_gen/user_part"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user_part.UserRegisterRequest) (resp *user_part.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user_part.UserLoginRequest) (resp *user_part.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}
