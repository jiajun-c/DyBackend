package rpc

import (
	"context"
	"errors"
	"tiktok/cmd/thumbup/kitex_gen/thumbuppart/thumbupservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var thumbupClient thumbupservice.Client

func initThumbupRPC() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString("etcd.addr")})
	logrus.Info("The etcd addr: ", viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	c, err := thumbupservice.NewClient(
		"thumbup_part",
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
	thumbupClient = c
}

func CommentAction(ctx context.Context, req *thumbuppart.comment_action_request) (*thumbuppart.comment_action_response, error) {
	resp, err := thumbupClient.GetCommentAction(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *thumbuppart.comment_list_request) (*thumbuppart.comment_list_response, error) {
	resp, err := thumbupClient.GetCommentList(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func FavoriteAction(ctx context.Context, req *thumbuppart.favorite_action_request) (*thumbuppart.favorite_action_response, error) {
	resp, err := thumbupClient.GetFavoriteAction(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func FavoriteList(ctx context.Context, req *thumbuppart.favorite_list_request) (*thumbuppart.favorite_list_response, error) {
	resp, err := thumbupClient.GetFavoriteList(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}
