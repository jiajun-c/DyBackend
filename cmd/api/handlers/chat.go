package handlers

import (
	context2 "context"
	"tiktok/cmd/api/rpc"
	"tiktok/internal/code"
	"tiktok/internal/errno"

	"github.com/gin-gonic/gin"
)

type MsgChatReq struct {
	FromUserId int64 `form:"from_user_id"`
	ToUserId   int64 `form:"to_user_id"`
}

type MsgActionReq struct {
	FromUserId int64  `form:"from_user_id"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
	Content    string `form:"content"`
}

func MessageChat(ctx *gin.Context) {
	var msgChatReq MsgChatReq
	err := ctx.BindQuery(&msgChatReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	resp, err := rpc.MessageChat(context2.Background(), &chatpart.douyin_message_chat_request{
		From_user_id: msgChatReq.FromUserId,
		To_user_id:   msgChatReq.ToUserId,
	})
	if err != nil {
		ctx.JSON(200, errno.MessageChatFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func MessageAction(ctx *gin.Context) {
	var msgActionReq MsgActionReq
	err := ctx.BindQuery(&msgActionReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	resp, err := rpc.MessageAction(context2.Background(), &chatpart.douyin_message_action_request{
		From_user_id: msgActionReq.FromUserId,
		To_user_id:   msgActionReq.ToUserId,
		Action_type:  msgActionReq.ActionType,
		Content:      msgActionReq.Content,
	})
	if err != nil {
		ctx.JSON(200, errno.MessageActionFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}
