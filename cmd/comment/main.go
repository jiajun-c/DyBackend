package main

import (
	"log"
	commentpart "tiktok/kitex_gen/commentpart/commentservice"
)

func main() {
	svr := commentpart.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
