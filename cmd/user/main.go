package main

import (
	"log"
	"tiktok/cmd/user/dal"
	user_part "tiktok/cmd/user/kitex_gen/user_part/userservice"
	"tiktok/internal/config"
)

func main() {
	config.Init("config_user.yaml")
	dal.Init()
	svr := user_part.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
