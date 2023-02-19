package service

import (
	"context"
	"tiktok/cmd/relation/dal/db"
	"tiktok/cmd/relation/pack"
	"tiktok/kitex_gen/relationpart"
)

type FollowingListService struct {
	ctx context.Context
}

func NewFollowingListService(ctx context.Context) *FollowingListService {
	return &FollowingListService{ctx: ctx}
}

func (s *FollowingListService) FollowingList(req *relationpart.DouyinRelationFollowListRequest) ([]*relationpart.User, error) {
	
	currentID := req.ActionUserId

	// TODO: userDB: check UserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("UserId not exist")
	// }

	followingLst, err := db.GetFollowingByUserID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	followingIDs := make([]int64, 0)
	for _, following := range followingLst {
		followingIDs = append(followingIDs, following.UserID)
	}
	// TODO: userDB：获取关注者信息
	// users, err := db.QueryUserByIds(s.ctx, followingIDs)
	// if err != nil {
	// 	return nil, err
	// }

	var relationMap map[int64]*db.Following
	if currentID == -1 {
		relationMap = nil
	} else {
		//查询当前用户与获取关注列表的关注关系
		relationMap, err = db.QueryFollowingRelation(s.ctx, currentID, followingIDs)
		if err != nil {
			return nil, err
		}
	}
	userList := pack.FollowUserList(currentID, users, relationMap)
	return userList, nil
}
