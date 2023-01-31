package router

import "github.com/gin-gonic/gin"

func Router(router *gin.Engine) {
	UserRouter(router)
}
