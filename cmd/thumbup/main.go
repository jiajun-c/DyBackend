package main

import (
	"log"
	thumbuppart "tiktok/cmd/thumbup/kitex_gen/thumbuppart/thumbupservice"
)

func main() {
	svr := thumbuppart.NewServer(new(ThumbupServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
