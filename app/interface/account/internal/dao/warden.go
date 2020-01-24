/*
@Time : 2020/1/24 13:43
@Author : Minus4
*/
package dao

import (
	pb "chat/app/interface/account/api"
	acc "chat/app/service/account/api"
	"context"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"
)

func NewWardenClient() (acc.AccountClient, error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 10),
		Timeout:           xtime.Duration(time.Millisecond * 500),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
		Zone:              env.Zone,
	}
	client, err := acc.NewClient(grpccfg)
	return client, err
}

func (d *dao) RawBasicInfo(ctx context.Context, uid int64) (info *pb.BasicInfo, err error) {
	resp, err := d.accClient.GetBasicInfo(ctx, &acc.BasicInfoReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}

	if resp == nil || resp.Uid == 0 {
		return nil, ecode.NothingFound
	}
	info = new(pb.BasicInfo)
	info.Email = resp.Email
	info.Uid = resp.Uid
	info.Nickname = resp.Nickname
	info.Sign = resp.Sign
	info.ProfilePicUrl = resp.ProfilePicUrl
	return
}
