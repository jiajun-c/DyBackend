package main

import (
	"context"
	videopart "tiktok/cmd/video/kitex_gen/videopart"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videopart.FeedRequest) (resp *videopart.FeedRresponse, err error) {
	// TODO: Your code here...
	return
}

// DoPublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DoPublishAction(ctx context.Context, req *videopart.DoPublishActionRequest) (resp *videopart.DoPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *videopart.GetPublishListRequest) (resp *videopart.GetPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
