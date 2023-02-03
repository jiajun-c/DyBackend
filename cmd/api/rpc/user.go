package rpc

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/userpart"
	"tiktok/kitex_gen/userpart/userservice"
	"time"
)

var userClient userservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString(viper.GetString("etcd.addr"))})
	logrus.Info("The etcd addr: ", viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		"userpart",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *userpart.UserRegisterRequest) error {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

func Login(ctx context.Context, req *userpart.UserLoginRequest) (*userpart.UserLoginResponse, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return nil, errors.New("error")
	}
	return resp, nil
}
