package service

import (
	"context"
	"errors"
	"favorite/dal/db"
	"favorite/dal/model"
	"favorite/kitex_gen/favoritepart"
)

type CreateFavoriteService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateFavoriteService(ctx context.Context) *CreateFavoriteService {
	return &CreateFavoriteService{ctx: ctx}
}

func (s *CreateFavoriteService) Favorite(req *favoritepart.FavoriteActionRequest) error {

	fav, err := db.QueryFavorite(s.ctx, 1, req.VideoId)
	if err != nil {
		return errors.New("error")
	}

	if len(fav) == 0 {
		db.DeleteFavorite(s.ctx, *fav[0])
	} else {
		db.CreateFavorite(s.ctx, []*model.Favorite{{UserID: 1, VideoID: req.VideoId}})
	}

	return errors.New("ok")

}
