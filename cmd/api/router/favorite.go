package router

import (
	"tiktok/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func FavoriteRouter(router *gin.Engine) {
	favoriteRouter := router.Group("/douyin/favorite")
	{
		favoriteRouter.POST("/action", handlers.Favorite)
	}
}
