package db

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	FollowerID int64     `json:"follower_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type Follower struct {
	gorm.Model
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	FollowerID   int64     `json:"follower_id"`
	FriendStatus bool      `json:"friend_status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

func (f *Following) TableName() string {
	return "following"
}

func (f *Follower) TableName() string {
	return "follower"
}

// QueryRelation
//
//	@Description: 		获取用户与用户切片关系(用户关注查询用户切片)
//	@param 	ctx			传入的上下文
//	@param 	userID		用户ID
//	@param 	queryID		查询用户ID切片
//	@return followings 	关注者切片
//	@return error		错误信息
func QueryFollowingRelation(ctx context.Context, userID int64, queryID []int64) (map[int64]*Following, error) {
	var followings []*Following
	err := DB.WithContext(ctx).Table("following").Where("follower_id = ? AND user_id IN ?", userID, queryID).Find(&followings).Error
	if err != nil {
		klog.Error("query following relation by ids " + err.Error())
		return nil, err
	}
	followingMap := make(map[int64]*Following)
	for _, relation := range followings {
		followingMap[relation.UserID] = relation
	}
	return followingMap, nil
}

// GetFollowerByUserID
//
//	@Description: 		获取用户粉丝列表
//	@param 	ctx			传入的上下文
//	@param 	userID		用户ID
//	@return []*Follower 	粉丝切片
//	@return error		错误信息
func GetFollowerByUserID(ctx context.Context, userID int64) ([]*Follower, error) {
	var followers []*Follower
	err := DB.WithContext(ctx).Table("follower").Where("user_id = ?", userID).Find(&followers).Error
	if err != nil {
		klog.Error("Query follower by user id failed " + err.Error())
		return nil, err
	}
	return followers, nil
}

// GetFollowingByUserID
//
//	@Description: 		获取用户关注列表
//	@param 	ctx			传入的上下文
//	@param 	userID		用户ID
//	@return []*Following 关注列表
//	@return error		错误信息
func GetFollowingByUserID(ctx context.Context, userID int64) ([]*Following, error) {
	var followings []*Following
	err := DB.WithContext(ctx).Table("following").Where("follower_id = ?", userID).Find(&followings).Error
	if err != nil {
		klog.Error("Query following list by user ID failed " + err.Error())
		return nil, err
	}
	return followings, nil
}

// GetFriendByUserID
//
//	@Description: 		获取用户朋友列表
//	@param 	ctx			传入的上下文
//	@param 	userID		用户ID
//	@return []*Follower 朋友列表
//	@return error		错误信息
func GetFriendByUserID(ctx context.Context, userID int64) ([]*Follower, error) {
	var friends []*Follower
	err := DB.WithContext(ctx).Table("following").Where("follower_id = ? AND friend_status = true", userID).Find(&friends).Error
	if err != nil {
		klog.Error("Query following list by user ID failed " + err.Error())
		return nil, err
	}
	return friends, nil
}

// Create
//
//	@Description: 		创建关注关系
//	@param 	ctx			传入的上下文
//	@param 	userID		被关注者ID
//	@param	followerID 	关注动作发起者ID
//	@return error		错误信息
func Create(ctx context.Context, userID int64, followerID int64) error {
	followingData := &Following{
		UserID:     userID,
		FollowerID: followerID,
	}
	followerData := &Follower{
		UserID:       userID,
		FollowerID:   followerID,
		FriendStatus: false,
	}
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 增加被关注者粉丝总数
		err := tx.Table("user").Where("id = ?", userID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			klog.Error("Relation action: follower_count increment failed " + err.Error())
			return err
		}
		// 增加关注动作发起者关注用户总数
		if err := tx.Table("user").Where("id = ?", followerID).Update("following_count", gorm.Expr("following_count + ?", 1)).Error; err != nil {
			klog.Error("Relation action: following_count increment failed " + err.Error())
			return err
		}
		// 查询被关注用户(userID)是否是当前用户(followerID)粉丝
		var follower *Follower
		if err := DB.WithContext(ctx).Table("follower").Where("user_id = ? AND follower_id = ?", followerID, userID).Find(&follower).Error; err != nil {
			klog.Error("Relation action: following relation query failed " + err.Error())
			return err
		}
		if follower != nil {
			followerData.FriendStatus = true
			// 更新被关注用户对当前用户的关注记录
			if err := DB.WithContext(ctx).Table("follower").Where("user_id = ? AND follower_id = ?", followerID, userID).Update("friend_status", true).Error; err != nil {
				klog.Error("Relation action: following account friend status update failed " + err.Error())
				return err
			}
		}

		// 创建关注记录
		if err := tx.Table("follower").Create(followerData).Error; err != nil {
			klog.Error("Relation action: create follower record falied " + err.Error())
			return err
		}
		if err := tx.Table("following").Create(followingData).Error; err != nil {
			klog.Error("Relation action: create following record falied " + err.Error())
			return err
		}

		return nil
	})
	return nil
}

// Delete
//
//	@Description: 		删除关注关系
//	@param 	ctx			传入的上下文
//	@param 	userID		被关注者ID
//	@param	followerID 	取消关注动作发起者ID
//	@return error		错误信息
func Delete(ctx context.Context, userID int64, followerID int64) error {
	var followerData *Follower
	var followingData *Following
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 减少被关注者粉丝总数
		err := tx.Table("user").Where("id = ?", userID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			klog.Error("Relation action: follower_count decrement failed " + err.Error())
			return err
		}
		// 减少取消关注动作发起者关注用户总数
		if err := tx.Table("user").Where("id = ?", followerID).Update("following_count", gorm.Expr("following_count - ?", 1)).Error; err != nil {
			klog.Error("Relation action: following_count decrement failed " + err.Error())
			return err
		}

		// 查询被关注用户(userID)是否是当前用户(followerID)粉丝
		var follower *Follower
		if err := DB.WithContext(ctx).Table("follower").Where("user_id = ? AND follower_id = ?", followerID, userID).Find(&follower).Error; err != nil {
			klog.Error("Relation action: following relation query failed " + err.Error())
			return err
		}
		if follower != nil {
			followerData.FriendStatus = true
			// 更新被关注用户对当前用户的关注记录
			if err := DB.WithContext(ctx).Table("follower").Where("user_id = ? AND follower_id = ?", followerID, userID).Update("friend_status", false).Error; err != nil {
				klog.Error("Relation action: following account friend status update failed " + err.Error())
				return err
			}
		}

		// 删除关注记录
		if err := tx.Table("follower").Where("user_id = ? AND follower_id = ?", userID, followerID).Delete(&followerData).Error; err != nil {
			klog.Error("Relation action: create follower record falied " + err.Error())
			return err
		}
		if err := tx.Table("following").Where("user_id = ? AND follower_id = ?", userID, followerID).Delete(&followingData).Error; err != nil {
			klog.Error("Relation action: create following record falied " + err.Error())
			return err
		}
		return nil
	})
	return nil
}
