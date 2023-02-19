// Package router
// @Description: 社交接口
package router

import (
	"tiktok/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func RelationRouter(router *gin.Engine) {
	relationRouter := router.Group("/relation")
	{
		relationRouter.POST("/action", handlers.RelationAction)
		relationRouter.GET("/follow/list", handlers.FollowList)
		relationRouter.GET("/follower/list", handlers.FollowerList)
		relationRouter.GET("/friend/list", handlers.FriendList)
	}
}
