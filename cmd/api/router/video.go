
package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/handlers"
)

func VideoRouter(router *gin.Engine) {
	feedRouter := router.GET("/feed", handlers.Feed)
	
	publishRouter := router.Group("/publish")
	{
		publishRouter.POST("/action", handlers.PublishAction)
		publishRouter.GET("/list", handlers.PublishList)
	}
}





