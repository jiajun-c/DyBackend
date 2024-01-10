package handlers

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/chat/kitex_gen/chatpart"
	"tiktok/internal/code"
	"tiktok/internal/errno"
	"tiktok/cmd/api/auth"
)

type RelationActionReq struct {
	Token		string `form:"token"`
	ToUserId    int64  `form:"to_user_id"`
	ActionType  int32  `form:"action_type"`
}

type FollowListReq struct {
	Token		string `form:"token"`
	UserId      int64  `form:"user_id"`
}

type FollowerListReq struct {
	Token		string `form:"token"`
	UserId      int64  `form:"user_id"`
}

type FriendListReq struct {
	Token		string `form:"token"`
	UserId      int64  `form:"user_id"`
}

func RelationAction(ctx *gin.Context) {
	var relAcReq RelationActionReq

	err := ctx.BindQuery(&relAcReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(relAcReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	response, err := rpc.RelationAction(context2.Background(), &relationpart.relation_action_request {
		To_user_id: relAcReq.ToUserID,
		Actio_type: relAcReq.ActionType,
	})
	if err!= nil {
		ctx.JSON(200, errno.RelationActionFailedErr)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}
}

func FollowList(ctx *gin.Context) {
	var followListReq FollowListReq

	err := ctx.BindQuery(&followListReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(followListReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	response, err := rpc.FollowList(context2.Background(), &relationpart.follow_list_request {
		User_id: followListReq.UserId,
	})
	if err!= nil {
		ctx.JSON(200, errno.FollowListFailedErr)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}
}

func FollowerList(ctx *gin.Context) {
	var followerListReq FollowListReq

	err := ctx.BindQuery(&followerListReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(followerListReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	response, err := rpc.FollowerList(context2.Background(), &relationpart.follower_list_request {
		User_id: followerListReq.UserId,
	})
	if err!= nil {
		ctx.JSON(200, errno.FollowerListFailedErr)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}
}

func FriendList(ctx *gin.Context) {
	var friendListReq FriendListReq

	err := ctx.BindQuery(&friendListReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(friendListReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	response, err := rpc.FriendList(context2.Background(), &relationpart.friend_list_request {
		User_id: friendListReq.UserId,
	})
	if err!= nil {
		ctx.JSON(200, errno.FriendListFailedErr)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}	
}

