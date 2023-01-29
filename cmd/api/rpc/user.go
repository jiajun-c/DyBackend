package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/spf13/viper"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/cmd/user/kitex_gen/userpart/userservice"
	"time"
)

var userClient userservice.Client

func initUserRpc() {
	c, err := userservice.NewClient(
		viper.GetString("ServiceName"),
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *userpart.UserRegisterRequest) error {

	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
