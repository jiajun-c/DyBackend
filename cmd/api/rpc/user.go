package rpc

import (
	"context"
	"errors"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/cmd/user/kitex_gen/userpart/userservice"
	"tiktok/internal/errno"
)

var userClient userservice.Client

func Register(ctx context.Context, req *userpart.UserRegisterRequest) error {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

func Login(ctx context.Context, req *userpart.UserLoginRequest) (*userpart.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return nil, errors.New("error")
	}
	return resp, nil
}
