// Code generated by Kitex v0.7.3. DO NOT EDIT.

package sample

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	dragon "kitexclient/kitex_gen/dragon"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DragonSay(ctx context.Context, req *dragon.DragonSayRequest, callOptions ...callopt.Option) (r *dragon.DragonSayResponse, err error)
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
	return &kSampleClient{
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

type kSampleClient struct {
	*kClient
}

func (p *kSampleClient) DragonSay(ctx context.Context, req *dragon.DragonSayRequest, callOptions ...callopt.Option) (r *dragon.DragonSayResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DragonSay(ctx, req)
}
