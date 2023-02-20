package main

import (
	favorite "favorite/kitex_gen/favoritepart/favoriteservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	// svr := demouser.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))
	// // svr := demouser.NewServer(new(UserServiceImpl))

	// err := svr.Run()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "favoriteservice"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
