package api

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

// AppID .
const AppID = "service.account"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (AccountClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewAccountClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
