package main

import (
	"fmt"
	"github.com/spf13/viper"
	"tiktok/internal/config"
	"time"
)

func main() {
	config.Init("config_video.yaml")
	fmt.Println(viper.GetString("test"))
	time.Sleep(5 * time.Second)
	fmt.Println(viper.GetString("test"))

}
