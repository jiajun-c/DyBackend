package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/cmd/api/router"
	"tiktok/cmd/api/rpc"
	"tiktok/internal/config"
	"tiktok/cmd/api/db"
)

func main() {
	config.Init("config_api.yaml")
	rpc.InitRPC()
	db.Init()
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
