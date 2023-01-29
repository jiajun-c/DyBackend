package main

import (
	"log"
	"tiktok/cmd/user/dal"
	user_part "tiktok/cmd/user/kitex_gen/user_part/userservice"
)

func main() {
	dal.Init()
	svr := user_part.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
