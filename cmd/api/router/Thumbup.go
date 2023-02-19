// Package router
// @Description: 点赞评论
package router

import (
	"tiktok/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func ThumbupRouter(router *gin.Engine) {
	commentRouter := router.Group("/comment")
	{
		commentRouter.POST("/action", handlers.CommentAction)
		commentRouter.GET("/list", handlers.CommentList)
	}

	favoriteRouter := router.Group("/favorite")
	{
		favoriteRouter.POST("/action", handlers.FavoriteAction)
		favoriteRouter.GET("/list", handlers.FavoriteList)
	}
}
