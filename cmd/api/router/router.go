package router

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {
	UserRouter(router)
	RelationRouter(router)
	VideoRouter(router)
	ThumbupRouter(router)
	CommentRouter(router)
	ChatRouter(router)
	FavoriteRouter(router)
}
