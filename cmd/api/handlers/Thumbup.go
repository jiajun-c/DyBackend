package handlers

import (
	context2 "context"
	"tiktok/cmd/api/rpc"
	"tiktok/internal/code"
	"tiktok/internal/errno"
	"tiktok/cmd/api/auth"

	"github.com/gin-gonic/gin"
)

type FavActionReq struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"`
}

type FavListReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type CommentActionReq struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentListReq struct {
	VideoId int64  `form:"video_id"`
	Token   string `form:"token"`
}

func CommentAction(ctx *gin.Context) {
	var commActionReq CommentActionReq
	err := ctx.BindQuery(&commActionReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(commActionReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.CommentAction(context2.Background(), &thumbuppart.comment_action_request{
		Video_id:     commActionReq.VideoId,
		Action_type:  commActionReq.ActionType,
		Comment_text: commActionReq.CommentText,
		Comment_id:   commActionReq.CommentId,
	})
	if err != nil {
		ctx.JSON(200, errno.CommentActionFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func CommentList(ctx *gin.Context) {
	var commListReq CommentListReq
	err := ctx.BindQuery(&commListReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(commListReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.CommentList(context2.Background(), &thumbuppart.comment_list_request{
		Video_id: commListReq.VideoId,
	})
	if err != nil {
		ctx.JSON(200, errno.CommentListFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func FavoriteAction(ctx *gin.Context) {
	var favActionReq FavActionReq
	err := ctx.BindQuery(&favActionReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(favActionReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.FavoriteAction(context2.Background(), &thumbuppart.favorite_action_request{
		Video_id:    favActionReq.VideoId,
		Action_type: favActionReq.ActionType,
	})
	if err != nil {
		ctx.JSON(200, errno.FavoriteActionFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func FavoriteList(ctx *gin.Context) {
	var favListReq FavListReq
	err := ctx.BindQuery(&favListReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(favListReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.FavoriteList(context2.Background(), &thumbuppart.favorite_list_request{
		User_id: favListReq.UserId,
	})
	if err != nil {
		ctx.JSON(200, errno.FavoriteListFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}




