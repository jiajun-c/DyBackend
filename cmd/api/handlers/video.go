package handlers

import (
	context2 "context"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/video/kitex_gen/videopart"
	"tiktok/internal/code"
	"tiktok/internal/errno"
	"tiktok/cmd/api/auth"

	"github.com/gin-gonic/gin"
)

type PublishActionReq struct {
	Token string `form:"token`
	Data  []byte `form:"data"`
	Title string `form:"title"`
}

type PublishListReq struct {
	UserId int64 `form:"user_id"`
	Token       string `form:"token"`
}

type FeedReq struct {
	Latest_time int64  `form:"latest_time"`
	Token       string `form:"token"`
}

func PublishAction(ctx *gin.Context) {
	var pubAcReq PublishActionReq
	err := ctx.BindQuery(&pubAcReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(pubAcReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.PublishAction(context2.Background(), &videopart.DoPublishActionRequest{
		Data:  pubAcReq.Data,
		Title: pubAcReq.Title,
	})
	if err != nil {
		ctx.JSON(200, errno.PublishActionFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func PublishList(ctx *gin.Context) {
	var pubListReq PublishListReq
	err := ctx.BindQuery(&pubListReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(pubListReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.PublishList(context2.Background(), &videopart.GetPublishListRequest{
		User_id: pubListReq.UserId,
	})
	if err != nil {
		ctx.JSON(200, errno.PublishListFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func Feed(ctx *gin.Context) {
	var feedReq FeedReq
	err := ctx.BindQuery(&feedReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	if !auth.Auth(feedReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.Feed(context2.Background(), &videopart.Feed{
		Latest_time: feedReq.Latest_time,
	})
	if err != nil {
		ctx.JSON(200, errno.PublishListFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}
