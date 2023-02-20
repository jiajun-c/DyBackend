package handlers

import (
	context2 "context"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/favorite/kitex_gen/favoritepart"
	"tiktok/internal/code"
	"tiktok/internal/errno"

	"github.com/gin-gonic/gin"
)

type FavoriteParam struct {
	Token      string `form:"token"`
	VideoID    int64  `form:"video_id"`
	ActionType int64  `form:"action_type"`
}

func Favorite(ctx *gin.Context) {
	var loginVar FavoriteParam
	err := ctx.BindQuery(&loginVar)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	resp, err := rpc.Favorite(context2.Background(), &favoritepart.FavoriteActionRequest{
		Token:   loginVar.Token,
		VideoID: loginVar.VideoID,
	})
	if err != nil {
		ctx.JSON(200, errno.AuthorizationFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}
