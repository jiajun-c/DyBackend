package handlers

import (
	context2 "context"
	"tiktok/cmd/api/rpc"
	"tiktok/internal/code"
	"tiktok/internal/errno"
	"tiktok/cmd/api/auth"

	"github.com/gin-gonic/gin"
)

type MsgActionRequest struct {
	Token		string `form:"token"`
	ToUserID 	string `form:"to_user_id"`
	ActionType  int32  `form:"action_type"`
	Content		string `form:"content"`
}

type MsgChatRequest struct { 
	Token	   string `form:"token"`
	PreMsgTime int64 `form:"pre_msg_time"`
	ToUserId   int64 `form:"to_user_id"`
}


func MessageAction(ctx *gin.Context) {
	var msgAcReq MsgActionRequest

	err := ctx.BindQuery(&msgAcReq)
	if err!= nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(msgAcReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	response, err := rpc.MessageAction(context2.Background(), &chatpart.douyin_message_action_request {
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

	if !auth.Auth(msgChReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	response, err := rpc.MessageChat(context2.Background(), &chatpart.douyin_message_chat_request {
		Pre_msg_time: msgChReq.PreMsgTime,
		To_user_id: msgChReq.ToUserID,
	})
	if err!= nil {
		ctx.JSON(200, errno.SystemError)
        return
	} else {
		ctx.JSON(code.StatusOK, response)
	}
}