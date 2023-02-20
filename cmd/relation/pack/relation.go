package pack

import (
	chatDB "tiktok/cmd/chat/dal/db"
	relationDB "tiktok/cmd/relation/dal/db"
	userDB "tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/relationpart"
)

func FollowUserList(currentId int64, users []*userDB.User, relationMap map[int64]*relationDB.Following) []*relationpart.User {
	userList := make([]*relationpart.User, 0)
	for _, user := range users {
		isFollow := false

		if currentId != -1 {
			_, ok := relationMap[int64(user.ID)]
			if ok {
				isFollow = true
			}
		}
		userList = append(userList, &relationpart.User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		})
	}
	return userList
}

func FriendUserList(currentId int64, users []*userDB.User, msgs []*chatDB.Message) []*relationpart.FriendUser {
	friendList := make([]*relationpart.FriendUser, 0)
	for i, user := range users {
		userRaw := &relationpart.User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
		msg := msgs[i]
		var msgType int64 = 0
		if msg.FromUserID == currentId {
			msgType = 1
		}
		friendList = append(friendList, &relationpart.FriendUser{
			UserInfo: userRaw,
			Message:  msg.Content,
			MsgType:  msgType,
		})
	}
	return friendList
}
