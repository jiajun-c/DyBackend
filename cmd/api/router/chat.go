package router

import (
	"tiktok/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func ChatRouter(router *gin.Engine) {
	messageRouter := router.Group("/message")
	{
		messageRouter.GET("/chat", handlers.MessageChat)
		messageRouter.POST("/action", handlers.MessageAction)
	}
}
