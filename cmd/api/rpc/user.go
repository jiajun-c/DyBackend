package rpc

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"tiktok/cmd/user/kitex_gen/userpart"
	"tiktok/cmd/user/kitex_gen/userpart/userservice"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString("etcd.addr")})
	logrus.Info("The etcd addr: ", viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		"user_part",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(300*time.Second),            // rpc timeout
		client.WithConnectTimeout(50000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *userpart.UserRegisterRequest) (*userpart.UserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *userpart.UserLoginRequest) (*userpart.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	logrus.Info(resp)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}
