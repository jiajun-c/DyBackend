package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/router"
	"tiktok/cmd/api/rpc"
	"tiktok/internal/config"
)

func main() {
	config.Init("config_api.yaml")
	rpc.InitRPC()
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
