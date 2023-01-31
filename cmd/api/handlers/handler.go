package handlers

import (
	"github.com/gin-gonic/gin"
	"tiktok/internal/code"
)

func WriteResponse(c *gin.Context) {
	c.JSON(code.StatusOK, gin.H{
		//"code":
	})
}
