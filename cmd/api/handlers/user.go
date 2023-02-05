package handlers

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/internal/code"
	"tiktok/internal/errno"
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

	resp, err := rpc.Login(context2.Background(), &userpart.UserLoginRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		ctx.JSON(200, errno.AuthorizationFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}

func Register(ctx *gin.Context) {
	var regVar UserParam
	err := ctx.BindQuery(&regVar)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
	}

	resp, err := rpc.Register(context2.Background(), &userpart.UserRegisterRequest{
		Username: regVar.UserName,
		Password: regVar.PassWord,
	})

	if err != nil {
		ctx.JSON(200, errno.UserRegisterFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}
