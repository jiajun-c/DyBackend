package service

import (
	"context"
	"tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/userpart"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *userpart.UserRegisterRequest) error {
	return db.CreateUser(s.ctx, []*db.User{{
		Name:     req.Username,
		Password: req.Password,
	}})
}
