package service

import (
	"context"
	"errors"
	"strconv"
	chatDB "tiktok/cmd/chat/dal/db"
	"tiktok/cmd/relation/config"
	"tiktok/cmd/relation/dal/db"
	"tiktok/cmd/relation/pack"
	"tiktok/cmd/relation/redis"
	"tiktok/kitex_gen/relationpart"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s *FriendListService) FriendList(req *relationpart.DouyinRelationFriendListRequest) ([]*relationpart.FriendUser, error) {
	currentID := req.ActionUserId
	if currentID != req.UserId {
		return nil, errors.New("No access to other user's friend list.")
	}
	friendIDs := make([]int64, 0)
	// TODO: userDB: check UserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("UserId not exist")
	// }

	userIDStr := strconv.Itoa(int(req.UserId))
	if cnt, _ := redis.RDBFriend.SCard(userIDStr).Result(); cnt != 0 {
		redis.RDBFriend.Expire(userIDStr, config.ExpireTime)
		friendIDStr, _ := redis.RDBFriend.SMembers(userIDStr).Result()
		for _, str := range friendIDStr {
			id, _ := strconv.Atoi(str)
			if id == -1 {
				continue
			}
			friendIDs = append(friendIDs, int64(id))
		}

	} else {

		friendLst, err := db.GetFriendByUserID(s.ctx, req.UserId)
		if err != nil {
			return nil, err
		}

		for _, friend := range friendLst {
			friendIDs = append(friendIDs, friend.FollowerID)
		}
		go redis.LoadFriendList(int(req.UserId), friendIDs)
	}

	// TODO: userDB：获取朋友信息
	users, err := db.QueryUserByIds(s.ctx, followerIDs)
	if err != nil {
		return nil, err
	}

	//TODO: chatDB: 双向聊天记录
	msgs := make([]*chatDB.Message, 0)
	for _, user := range users {
		msg, err := chatDB.GetChatHistory(s.ctx, req.ActionUserId, user.ID)
		if err != nil || msg == nil {
			msgs = append(msgs, nil)
		}
		msgs = append(msgs, msg[0])
	}

	friendList := pack.FriendUserList(currentID, users, msgs)
	return friendList, nil

}
