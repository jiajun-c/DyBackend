package rpc

import (
	"context"
	"errors"
	"tiktok/cmd/chat/kitex_gen/chatpart/chatservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var chatClient Chatservice.Client

func initChatRPC() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString("etcd.addr")})
	logrus.Info("The etcd addr: ", viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	c, err := chatservice.NewClient(
		"chat_part",
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
	chatClient = c
}

func MessageChat(ctx context.Context, req *chatpart.douyin_message_chat_request) (*chatpart.douyin_message_chat_response, error) {
	resp, err := chatClient.GetChatHistory(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func MessageAction(ctx context.Context, req *chatpart.douyin_message_action_request) (*chatpart.douyin_message_action_response, error) {
	resp, err := chatClient.DoMessageAction(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}
