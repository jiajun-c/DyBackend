package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/rpc"
	"tiktok/internal/code"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/userpart"
)

type UserParam struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

func Login(ctx *gin.Context) {
	var loginVar UserParam
	err := ctx.BindQuery(&loginVar)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
		return
	}

	resp, err := rpc.Login(ctx, &userpart.UserLoginRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		ctx.JSON(200, errno.AuthorizationFailedErr)
	}
	ctx.JSON(code.StatusOK, resp)
}
