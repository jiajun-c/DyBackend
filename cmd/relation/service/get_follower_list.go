package service

import (
	"context"
	"tiktok/cmd/relation/dal/db"
	"tiktok/cmd/relation/pack"
	"tiktok/kitex_gen/relationpart"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

func (s *FollowerListService) FollowerList(req *relationpart.DouyinRelationFollowerListRequest) ([]*relationpart.User, error) {
	
	currentID := req.ActionUserId

	// TODO: userDB: check UserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("UserId not exist")
	// }

	followerLst, err := db.GetFollowerByUserID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	followerIDs := make([]int64, 0)
	for _, follower := range followerLst {
		followerIDs = append(followerIDs, follower.FollowerID)
	}
	// TODO: userDB：获取粉丝信息
	// users, err := db.QueryUserByIds(s.ctx, followerIDs)
	// if err != nil {
	// 	return nil, err
	// }

	var relationMap map[int64]*db.Following
	if currentID == -1 {
		relationMap = nil
	} else {
		//查询当前用户与获取粉丝列表的关注关系
		relationMap, err = db.QueryFollowingRelation(s.ctx, currentID, followerIDs)
		if err != nil {
			return nil, err
		}
	}
	userList := pack.FollowUserList(currentID, users, relationMap)
	return userList, nil

}
