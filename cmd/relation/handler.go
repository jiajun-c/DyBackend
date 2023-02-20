package main

import (
	"context"
	"tiktok/cmd/relation/service"
	"tiktok/internal/app"
	"tiktok/internal/errno"
	relationpart "tiktok/kitex_gen/relationpart"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relationpart.DouyinRelationActionRequest) (resp *relationpart.DouyinRelationActionResponse, err error) {
	resp = new(relationpart.DouyinRelationActionResponse)

	if req.UserId <= 0 || req.ToUserId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = app.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	if err = service.NewRelationActionService(ctx).RelationAction(req); err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = app.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relationpart.DouyinRelationFollowListRequest) (resp *relationpart.DouyinRelationFollowListResponse, err error) {
	resp = new(relationpart.DouyinRelationFollowListResponse)

	if req.ActionUserId < 0 || req.UserId <= 0 {
		resp.BaseResp = app.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	followingLst, err := service.NewFollowingListService(ctx).FollowingList(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserList = followingLst
	resp.BaseResp = app.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relationpart.DouyinRelationFollowerListRequest) (resp *relationpart.DouyinRelationFollowerListResponse, err error) {
	resp = new(relationpart.DouyinRelationFollowerListResponse)

	if req.ActionUserId < 0 || req.UserId <= 0 {
		resp.BaseResp = app.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	followerLst, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserList = followerLst
	resp.BaseResp = app.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relationpart.DouyinRelationFriendListRequest) (resp *relationpart.DouyinRelationFriendListResponse, err error) {
	resp = new(relationpart.DouyinRelationFriendListResponse)

	if req.ActionUserId < 0 || req.UserId <= 0 {
		resp.BaseResp = app.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	friendLst, err := service.NewFriendListService(ctx).FriendList(req)
	if err != nil {
		resp.BaseResp = app.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserList = friendLst
	resp.BaseResp = app.BuildBaseResp(errno.Success)
	return resp, nil
}
