package main

import (
	"log"
	userpart "tiktok/cmd/user/kitex_gen/userpart/userservice"
)

func main() {
	svr := userpart.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
