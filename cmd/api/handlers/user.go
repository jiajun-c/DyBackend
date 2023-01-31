package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/internal/code"
	"tiktok/internal/errno"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func Login(ctx *gin.Context) {
	var loginVar UserParam
	if err := ctx.Bind(&loginVar); err != nil {
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
