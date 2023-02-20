// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	videopart "tiktok/cmd/thumbup/kitex_gen/videopart"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, Req *videopart.FeedRequest, callOptions ...callopt.Option) (r *videopart.FeedRresponse, err error)
	DoPublishAction(ctx context.Context, Req *videopart.DoPublishActionRequest, callOptions ...callopt.Option) (r *videopart.DoPublishActionResponse, err error)
	GetPublishList(ctx context.Context, Req *videopart.GetPublishListRequest, callOptions ...callopt.Option) (r *videopart.GetPublishListResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) Feed(ctx context.Context, Req *videopart.FeedRequest, callOptions ...callopt.Option) (r *videopart.FeedRresponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}

func (p *kVideoServiceClient) DoPublishAction(ctx context.Context, Req *videopart.DoPublishActionRequest, callOptions ...callopt.Option) (r *videopart.DoPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoPublishAction(ctx, Req)
}

func (p *kVideoServiceClient) GetPublishList(ctx context.Context, Req *videopart.GetPublishListRequest, callOptions ...callopt.Option) (r *videopart.GetPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishList(ctx, Req)
}
