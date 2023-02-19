package handlers

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/user/kitex_gen/chatpart"
	"tiktok/internal/code"
	"tiktok/internal/errno"
)

type MsgActionRequest struct {
	FromUserId	int64  `form:"from_user_id"`
	ToUserID 	string `form:"to_user_id"`
	ActionType  int32  `form:"action_type"`
	Content		string `form:"content"`
}

type MsgChatRequest struct { 
	FromUserId int64 `form:"from_user_id"`
	ToUserId   int64 `form:"to_user_id"`
}

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


func   (ctx *gin.Context) {
	var msgAcReq MsgActionRequest

	err := ctx.BindQuery(&msgAcReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	response, err := rpc.MessageAction(context2.Background(), &chatpart.douyin_message_action_request {
		From_user_id: msgAcReq.FromUserId,
		To_user_id: msgAcReq.ToUserID,
		Action_type: msgAcReq.ActionType,
		Content: msgAcReq.Content,
	})
	if err!= nil {
		ctx.JSON(200, errno.SystemError)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}
}

func MessageChat(ctx *gin.Context) {
	var msgChReq MsgChatRequest

	err := ctx.BindQuery(&msgChReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	response, err := rpc.MessageChat(context2.Background(), &chatpart.douyin_message_chat_request {
		From_user_id: msgChReq.FromUserId,
		To_user_id: msgChReq.ToUserID,
	})
	if err!= nil {
		ctx.JSON(200, errno.SystemError)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}
}

func RelationAction(ctx *gin.Context) {
	var relAcReq RelationActionReq

	err := ctx.BindQuery(&relAcReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	response, err := rpc.RelationAction(context2.Background(), &relationpart.relation_action_request {
		Token: relAcReq.Token,
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

	response, err := rpc.FollowList(context2.Background(), &relationpart.follow_list_request {
		Token: followListReq.Token,
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

	response, err := rpc.MessageChat(context2.Background(), &relationpart.follower_list_request {
		Token: followerListReq.Token,
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

	response, err := rpc.FriendList(context2.Background(), &relationpart.friend_list_request {
		Token: friendListReq.Token,
		User_id: friendListReq.UserId,
	})
	if err!= nil {
		ctx.JSON(200, errno.FriendListFailedErr)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}	
}

