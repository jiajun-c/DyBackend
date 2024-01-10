package handlers

import (
	context2 "context"
	"tiktok/cmd/api/rpc"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/internal/code"
	"tiktok/internal/errno"
	"tiktok/cmd/api/db"
	"tiktok/cmd/api/auth"

	"github.com/gin-gonic/gin"
)

type UserParam struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserInfoReq struct {
	User_id int64  `form:"user_id"`
	Token   string `form:"token"`
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
		u, err := db.GetUserByUsername(loginVar.Username)
		if err != nil {
			ctx.JSON(200, errno.ParamErr)
			return
		}

		resp.UserId = u.Uid
		resp.Token = auth.GenerateToken(loginVar.Username)
		ctx.JSON(code.StatusOK, resp)
	}
	return
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
		u, err := db.GetUserByUsername(loginVar.Username)
		if err != nil {
			ctx.JSON(200, errno.ParamErr)
			return
		}
		
		resp.UserId = u.Uid
		resp.Token = auth.GenerateToken(loginVar.Username)
		ctx.JSON(code.StatusOK, resp)
	}
	return
}

func Info(ctx *gin.Context) {
	var usrInfoReq UserInfoReq
	err := ctx.BindQuery(&usrInfoReq)
	if err != nil {
		ctx.JSON(200, errno.ParamErr)
	}

	if !auth.Auth(usrInfoReq.Token) {
		ctx.JSON(200, errno.TokenFailedErr)
		return
	}

	resp, err := rpc.Info(context2.Background(), &userpart.UserInfoRequest{
		User_id: usrInfoReq.User_id,
	})

	if err != nil {
		ctx.JSON(200, errno.UserInfoFailedErr)
	} else {
		ctx.JSON(code.StatusOK, resp)
	}
}
