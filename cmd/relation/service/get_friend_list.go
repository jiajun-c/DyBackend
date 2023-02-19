package service

import (
	"context"
	"errors"
	chatDB "tiktok/cmd/chat/dal/db"
	"tiktok/cmd/relation/dal/db"
	"tiktok/cmd/relation/pack"
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

	// TODO: userDB: check UserID existence
	// users, err := db.QueryUserByIds(s.ctx, []int64{req.ToUserId})
	// if err != nil {
	// 	return nil, err
	// }
	// if len(users) == 0 {
	// 	return nil, errors.New("UserId not exist")
	// }

	friendLst, err := db.GetFriendByUserID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	friendIDs := make([]int64, 0)
	for _, friend := range friendLst {
		friendIDs = append(friendIDs, friend.FollowerID)
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
