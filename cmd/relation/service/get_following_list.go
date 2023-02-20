package service

import (
	"context"
	"strconv"
	"tiktok/cmd/relation/config"
	"tiktok/cmd/relation/dal/db"
	"tiktok/cmd/relation/pack"
	"tiktok/cmd/relation/redis"
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

	followingIDs := make([]int64, 0)

	// TODO: userDB: check UserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("UserId not exist")
	// }

	followerIDStr := strconv.Itoa(int(req.UserId))
	if cnt, _ := redis.RDBFollowing.SCard(followerIDStr).Result(); cnt != 0 {
		redis.RDBFollowing.Expire(followerIDStr, config.ExpireTime)
		followingIDStr, _ := redis.RDBFollowing.SMembers(followerIDStr).Result()
		for _, str := range followingIDStr {
			id, _ := strconv.Atoi(str)
			if id == -1 {
				continue
			}
			followingIDs = append(followingIDs, int64(id))
		}
	} else {
		followingLst, err := db.GetFollowingByUserID(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}

		for _, following := range followingLst {
			followingIDs = append(followingIDs, following.UserID)
		}
		go redis.LoadFollowingList(int(req.UserId), followingIDs)
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
