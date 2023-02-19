// Package router
// @Description: 评论接口
package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/handlers"
)

func CommentRouter(router *gin.Engine) {
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("/action", handlers.CommentAction)
		commentRouter.GET("/list", handlers.CommentList)
	}
}