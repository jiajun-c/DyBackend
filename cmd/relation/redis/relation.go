package redis

import (
	"strconv"
	"tiktok/cmd/relation/config"
)

// Follow
//
//	@Description: 		对redis数据库注入新增关注关系
//	@param 	userID		被关注者ID
//	@param	followerID 	关注动作发起者ID
//	@return error		错误信息
func Follow(userID int, followerID int) error {

	userIDStr := strconv.Itoa(int(userID))
	followerIDStr := strconv.Itoa(int(followerID))

	// FollowerDB 存在被关注者的粉丝Set
	if cnt, _ := RDBFollower.SCard(userIDStr).Result(); cnt != 0 {
		if err := RDBFollower.SRem(userIDStr, followerID).Err(); err != nil {
			return err
		}
		RDBFollower.Expire(userIDStr, config.ExpireTime)
	}

	// FollowingDB存在关注动作发起者的关注Set
	if cnt, _ := RDBFollowing.SCard(followerIDStr).Result(); cnt != 0 {
		if err := RDBFollowing.SRem(followerIDStr, userID).Err(); err != nil {
			return err
		}
		RDBFollowing.Expire(followerIDStr, config.ExpireTime)
	}
	// FriendDB是否存在两人的朋友列表
	if cnt, _ := RDBFriend.SCard(followerIDStr).Result(); cnt != 0 {
		RDBFriend.SRem(followerIDStr, userID)
		RDBFriend.Expire(followerIDStr, config.ExpireTime)
	}
	if cnt, _ := RDBFriend.SCard(userIDStr).Result(); cnt != 0 {
		RDBFriend.SRem(userIDStr, followerID)
		RDBFriend.Expire(userIDStr, config.ExpireTime)
	}
	return nil
}

// UnFollow
//
//	@Description: 		删除redis数据库内的关注关系
//	@param 	userID		被关注者ID
//	@param	followerID 	取关动作发起者ID
//	@return error		错误信息
func UnFollow(userID int, followerID int) error {

	userIDStr := strconv.Itoa(int(userID))
	followerIDStr := strconv.Itoa(int(followerID))

	// FollowerDB 存在被关注者的粉丝Set
	if cnt, _ := RDBFollower.SCard(userIDStr).Result(); cnt != 0 {
		if err := RDBFollower.SAdd(userIDStr, followerID).Err(); err != nil {
			return err
		}
		RDBFollower.Expire(userIDStr, config.ExpireTime)
	}

	// FollowingDB存在关注动作发起者的关注Set
	if cnt, _ := RDBFollowing.SCard(followerIDStr).Result(); cnt != 0 {
		if err := RDBFollowing.SAdd(followerIDStr, userID).Err(); err != nil {
			return err
		}
		RDBFollowing.Expire(followerIDStr, config.ExpireTime)
	}
	// FriendDB是否存在两人的朋友列表
	if cnt, _ := RDBFriend.SCard(followerIDStr).Result(); cnt != 0 {
		RDBFriend.Del(followerIDStr)
	}
	if cnt, _ := RDBFriend.SCard(userIDStr).Result(); cnt != 0 {
		RDBFriend.Del(userIDStr)
	}
	return nil
}

// LoadFollowerList
//
//	@Description: 		将粉丝列表存入redis数据库
//	@param 	userID		被关注者ID
//	@param	followerID 	粉丝ID切片
func LoadFollowerList(userID int, followerID []int64) {

	// 初始化该用户粉丝set，加入-1防止脏读
	userIdStr := strconv.Itoa(int(userID))
	RDBFollowing.SAdd(userIdStr, -1)

	for _, id := range followerID {
		RDBFollowing.SAdd(userIdStr, id)
	}
	RDBFollowing.Expire(userIdStr, config.ExpireTime)
}

// LoadFollowingList
//
//	@Description: 		将关注列表存入redis数据库
//	@param	followerID 	粉丝ID
//	@param 	userID		被关注者ID切片
func LoadFollowingList(followerID int, userID []int64) {

	// 初始化该用户关注set，加入-1防止脏读
	followerIdStr := strconv.Itoa(int(followerID))
	RDBFollowing.SAdd(followerIdStr, -1)

	for _, id := range userID {
		RDBFollowing.SAdd(followerIdStr, id)
	}
	RDBFollowing.Expire(followerIdStr, config.ExpireTime)
}

// LoadFriendList
//
//	@Description: 		将朋友列表存入redis数据库
//	@param	userID 		用户ID
//	@param 	friendID	朋友ID切片
func LoadFriendList(userID int, friendID []int64) {

	// 初始化该用户关注set，加入-1防止脏读
	friendIdStr := strconv.Itoa(int(userID))
	RDBFriend.SAdd(friendIdStr, -1)

	for _, id := range friendID {
		RDBFriend.SAdd(friendIdStr, id)
	}
	RDBFriend.Expire(friendIdStr, config.ExpireTime)
}
