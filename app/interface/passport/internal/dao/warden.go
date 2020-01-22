/*
@Time : 2020/1/21 17:40
@Author : Minus4
*/
package dao

import (
	"chat/app/interface/passport/internal/model"
	acc "chat/app/service/account/api"
	"context"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"
)

func NewWardenClient() (acc.AccountClient, error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 10),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
		Zone:              env.Zone,
	}
	client, err := acc.NewClient(grpccfg)
	return client, err
}

func (d *dao) Register(ctx context.Context, email, password string) (uid int64, err error) {
	registerRequest := &acc.RegisterReq{
		Email:    email,
		Password: password,
	}
	remoteResp, err := d.accClient.Register(ctx, registerRequest)
	if err != nil {
		return 0, err
	}
	uid = remoteResp.Uid
	_ = d.DeleteCacheAuthInfo(ctx, email)
	return
}

func (d *dao) RawAuthInfo(ctx context.Context, email string) (info *model.AuthInfo, err error) {
	authReq := &acc.AuthReq{
		Email: email,
	}
	authResp, err := d.accClient.GetAuthInfo(ctx, authReq)
	if err != nil {
		return nil, err
	}
	info = new(model.AuthInfo)
	info.Sum = authResp.Password
	info.Uid = authResp.Uid
	return
}
