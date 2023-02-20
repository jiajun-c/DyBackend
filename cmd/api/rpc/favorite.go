package rpc

import (
	"context"
	"errors"
	"tiktok/cmd/favorite/kitex_gen/favoritepart"
	"tiktok/cmd/favorite/kitex_gen/favoritepart/favoriteservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var favoriteClient favoriteservice.Client

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // r不应重复使用。
	if err != nil {
		log.Fatal(err)
	}
	client, err := favoriteservice.NewClient("favoriteservice", client.WithResolver(r))
	favoriteClient = client
}

func Favorite(ctx context.Context, req *favoritepart.FavoriteActionRequest) (*favoritepart.FavoriteActionResponse, error) {
	resp, err := favoriteClient.Favorite(ctx, req)
	if err != nil || resp.StatusCode != 0 {
		return resp, errors.New("error")
	}
	return resp, nil
}
