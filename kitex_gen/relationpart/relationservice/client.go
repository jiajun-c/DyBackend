// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relationpart "tiktok/kitex_gen/relationpart"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, Req *relationpart.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationActionResponse, err error)
	GetFollowList(ctx context.Context, Req *relationpart.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationFollowListResponse, err error)
	GetFollowerList(ctx context.Context, Req *relationpart.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationFollowerListResponse, err error)
	GetFriendList(ctx context.Context, Req *relationpart.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationFriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, Req *relationpart.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, Req)
}

func (p *kRelationServiceClient) GetFollowList(ctx context.Context, Req *relationpart.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, Req)
}

func (p *kRelationServiceClient) GetFollowerList(ctx context.Context, Req *relationpart.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, Req)
}

func (p *kRelationServiceClient) GetFriendList(ctx context.Context, Req *relationpart.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *relationpart.DouyinRelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, Req)
}
