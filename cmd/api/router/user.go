package router

import (
	"tiktok/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handlers.Register)
		userRouter.POST("/login", handlers.Login)
		userRouter.GET("/", handlers.Info)
	}
}
