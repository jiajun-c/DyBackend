package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"tiktok/cmd/chat/dal"
	"tiktok/internal/config"

	"tiktok/kitex_gen/chatpart/chatservice"
)

func main() {
	config.Init("config_chat.yaml")
	dal.Init()
	r, err := etcd.NewEtcdRegistry([]string{viper.GetString("etcd.addr")})
	logrus.Info("etcd addr:" + viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:18889")
	if err != nil {
		panic(err)
	}
	svr := chatservice.NewServer(new(ChatServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "chat_part"}), // server name
		server.WithServiceAddr(addr),                                                     // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),               // limit
		//server.WithMuxTransport(),                                                        // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),
		//server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler // tracer
		server.WithRegistry(r), // registry
	)
	logrus.Info("start the svr")
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
