package rpc

import (
	"context"
	"errors"
	"tiktok/cmd/relation/kitex_gen/relationpart/relationservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var relationClient relationservice.Client

func initRelationRPC() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString("etcd.addr")})
	logrus.Info("The etcd addr: ", viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	c, err := relationservice.NewClient(
		"relation_part",
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
	relationClient = c
}

func RelationAction(ctx context.Context, req *relationpart.relation_action_request) (*relationpart.relation_action_response, error) {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func FollowList(ctx context.Context, req *relationpart.follow_list_request) (*relationpart.follow_list_response, error) {
	resp, err := relationClient.FollowList(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func FollowerList(ctx context.Context, req *relationpart.follower_list_request) (*relationpart.follower_list_response, error) {
	resp, err := relationClient.FollowerList(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func FriendList(ctx context.Context, req *relationpart.friend_list_request) (*relationpart.friend_list_response, error) {
	resp, err := relationClient.FriendList(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}
