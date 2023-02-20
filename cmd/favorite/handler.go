package main

import (
	"context"
	favoritepart "favorite/kitex_gen/favoritepart"
	"favorite/service"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// Favorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Favorite(ctx context.Context, req *favoritepart.FavoriteActionRequest) (resp *favoritepart.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	err = service.NewCreateFavoriteService(ctx).Favorite(req)
	resp = favoritepart.NewFavoriteActionResponse()
	ss := err.Error()
	resp.SetStatusMsg(&ss)
	return resp, nil
}
