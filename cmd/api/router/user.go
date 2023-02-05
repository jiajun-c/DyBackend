package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/handlers"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handlers.Register)
		userRouter.POST("/login", handlers.Login)
	}
}
