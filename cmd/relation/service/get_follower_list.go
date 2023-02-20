package service

import (
	"context"
	"strconv"
	"tiktok/cmd/relation/config"
	"tiktok/cmd/relation/dal/db"
	"tiktok/cmd/relation/pack"
	"tiktok/cmd/relation/redis"
	userDB "tiktok/cmd/user/dal/db"
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
	followerIDs := make([]int64, 0)

	// TODO: userDB: check UserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("UserId not exist")
	// }
	// refer in Redis
	userIDStr := strconv.Itoa(int(req.UserId))
	if cnt, _ := redis.RDBFollower.SCard(userIDStr).Result(); cnt != 0 {
		redis.RDBFollower.Expire(userIDStr, config.ExpireTime)
		followerIDStr, _ := redis.RDBFollower.SMembers(userIDStr).Result()
		for _, str := range followerIDStr {
			id, _ := strconv.Atoi(str)
			if id == -1 {
				continue
			}
			followerIDs = append(followerIDs, int64(id))
		}

	} else {
		followerLst, err := db.GetFollowerByUserID(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}

		for _, follower := range followerLst {
			followerIDs = append(followerIDs, follower.FollowerID)
		}
		go redis.LoadFollowerList(int(req.UserId), followerIDs)
	}

	// TODO: userDB：获取粉丝信息
	users, err := userDB.QueryUserByIds(s.ctx, followerIDs)
	if err != nil {
		return nil, err
	}

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
