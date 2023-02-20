package service

import (
	"context"
	"errors"
	"tiktok/cmd/relation/constants"
	"tiktok/cmd/relation/dal/db"
	"tiktok/kitex_gen/relationpart"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relationpart.DouyinRelationActionRequest) error {
	currentID := req.UserId

	// TODO: check ToUserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return err
	// }
	// if len(users) == 0 {
	// 	return errors.New("toUserId not exist")
	// }

	if req.ActionType == constants.Follow {
		if err := db.Create(s.ctx, currentID, req.ToUserId); err != nil {
			return err
		}
		return nil
	}
	if req.ActionType == constants.UnFollow {
		if err := db.Delete(s.ctx, currentID, req.ToUserId); err != nil {
			return err
		}
		return nil
	}
	return errors.New("RelationAction Service Unknown Error.")
}
