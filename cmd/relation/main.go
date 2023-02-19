package main

import (
	"log"
	relationpart "tiktok/kitex_gen/relationpart/relationservice"
)

func main() {
	svr := relationpart.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
