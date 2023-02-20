package main

import (
	"context"
	thumbuppart "tiktok/cmd/thumbup/kitex_gen/thumbuppart"
)

// ThumbupServiceImpl implements the last service interface defined in the IDL.
type ThumbupServiceImpl struct{}

// GetFavoriteAction implements the ThumbupServiceImpl interface.
func (s *ThumbupServiceImpl) GetFavoriteAction(ctx context.Context, req *thumbuppart.FavoriteActionRequest) (resp *thumbuppart.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteList implements the ThumbupServiceImpl interface.
func (s *ThumbupServiceImpl) GetFavoriteList(ctx context.Context, req *thumbuppart.FavoriteListRequest) (resp *thumbuppart.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCommentAction implements the ThumbupServiceImpl interface.
func (s *ThumbupServiceImpl) GetCommentAction(ctx context.Context, req *thumbuppart.CommentActionRequest) (resp *thumbuppart.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCommentList implements the ThumbupServiceImpl interface.
func (s *ThumbupServiceImpl) GetCommentList(ctx context.Context, req *thumbuppart.CommentListRequest) (resp *thumbuppart.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
