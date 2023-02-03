package service

import (
	"context"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/kitex_gen/userpart"
)

type LoginUserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{ctx: ctx}
}

func (s *LoginUserService) LoginUser(req *userpart.UserLoginRequest) bool {
	return db.CheckUser(s.ctx, req.Username, req.Password)
}
