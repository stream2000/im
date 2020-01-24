/*
@Time : 2020/1/24 17:14
@Author : Minus4
*/
package main

import (
	"chat/app/service/account/api"
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"
)

func NewWardenClient() (api.AccountClient, error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 10),
		Timeout:           xtime.Duration(time.Second * 500),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
		Zone:              env.Zone,
	}
	client, err := api.NewClient(grpccfg)
	return client, err
}

func main() {
	c, err := NewWardenClient()
	if err != nil {
		panic(err)
	}
	info, err := c.GetBasicInfo(context.Background(), &api.BasicInfoReq{
		Uid: 1,
	})
	if err != nil {
		panic(err)
	}
	if info.Uid == 0 {
		fmt.Println("Empty")
	}
}
