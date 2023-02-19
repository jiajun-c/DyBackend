package service

import (
	"context"
	"errors"
	"tiktok/cmd/comment/dal/db"
	"tiktok/cmd/comment/pack"
	userdb "tiktok/cmd/user/dal/db"

	// videodb "tiktok/cmd/video/dal/db"
	// relationdb "tiktok/cmd/relation/dal/db"
	"tiktok/kitex_gen/commentpart"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService new CommentListService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *commentpart.CommentListRequest) ([]*commentpart.Comment, error) {
	// 验证视频Id是否存在
	videos, err := videodb.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("video not exist")
	}

	// 根据视频id获取一组评论
	comments, err := db.QueryCommentByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	// 获取该组评论的用户id
	userIds := make([]int64, 0)
	for _, comment := range comments {
		userIds = append(userIds, comment.UserId)
	}
	// 获取一系列用户信息
	users, err := userdb.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*userdb.User)
	for _, user := range users {
		userMap[int64(user.ID)] = user
	}

	var relationMap map[int64]*relationdb.Relation
	if req.UserId == -1 {
		relationMap = nil
	} else {
		// 获取一系列关注信息
		relationMap, err = relationdb.QueryRelationByIds(s.ctx, req.UserId, userIds)
		if err != nil {
			return nil, err
		}
	}

	commentList := pack.CommentList(req.UserId, comments, userMap, relationMap)
	return commentList, nil
}
