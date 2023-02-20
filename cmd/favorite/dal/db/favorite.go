package db

import (
	"context"
	"favorite/dal/model"
)

func QueryFavorite(ctx context.Context, userId int64, videoId int64) ([]*model.Favorite, error) {
	return Q.WithContext(ctx).Favorite.Where(Q.Favorite.UserID.Eq(userId)).Where(Q.Favorite.VideoID.Eq(videoId)).Find()
}

func CreateFavorite(ctx context.Context, favorite []*model.Favorite) error {
	return Q.WithContext(ctx).Favorite.Create(favorite...)
}

func DeleteFavorite(ctx context.Context, favorite model.Favorite) {
	Q.WithContext(ctx).Favorite.Delete(&favorite)
}
