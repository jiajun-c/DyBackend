package rpc

import (
	"context"
	"errors"
	"tiktok/cmd/video/kitex_gen/videopart"
	"tiktok/cmd/video/kitex_gen/videopart/videoservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var videoClient videoservice.Client

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{viper.GetString("etcd.addr")})
	logrus.Info("The etcd addr: ", viper.GetString("etcd.addr"))
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		"video_part",
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
	videoClient = c
}

func PublishAction(ctx context.Context, req *videopart.DoPublishActionRequest) (*videopart.DoPublishActionResponse, error) {
	resp, err := videoClient.DoPublishAction(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func PublishList(ctx context.Context, req *videopart.GetPublishListRequest) (*videopart.GetPublishListResponse, error) {
	resp, err := videoClient.GetPublishList(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}

func Feed(ctx context.Context, req *videopart.FeedRequest) (*videopart.FeedResponse, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}
