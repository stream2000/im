/*
@Time : 2020/1/24 23:55
@Author : Minus4
*/
package dao

import (
	grp "chat/app/service/group/api"
	"context"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/pkg/errors"
	"time"
)

func NewWardenClient() (grp.GroupClient, error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 35),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
		Zone:              env.Zone,
	}
	client, err := grp.NewClient(grpccfg)
	return client, err
}

// AllGroups return all groups in the system without paging
// TODO store certain group info in redis, and use map to update the group info
func (d *dao) AllGroups(ctx context.Context) (resp *grp.AllGroups, err error) {
	groups, err := d.grpClient.GetAllGroups(ctx, &grp.GetAllGroupsReq{
		Uid: 0,
	})
	if err != nil {
		return
	}
	if groups == nil {
		return nil, nil
	}
	return groups, nil
}

func (d *dao) AllGroupsByUserId(ctx context.Context, uid int64) (resp *grp.AllGroups, err error) {
	groups, err := d.grpClient.GetAllGroups(ctx, &grp.GetAllGroupsReq{
		Uid: uid,
	})
	resp = new(grp.AllGroups)
	if err != nil {
		return resp, errors.Wrapf(err, "%d", uid)
	}
	if groups == nil || len(groups.Groups) == 0 {
		return resp, nil
	}
	resp.Groups = groups.Groups
	return
}

func (d *dao) GerGroupDetailedInfo(ctx context.Context, gid int64) (resp *grp.GroupInfo, err error) {
	resp, err = d.grpClient.GetGroupInfo(ctx, &grp.GroupInfoByIdReq{
		Gid: gid,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "%d", gid)
	}
	// solve the nil value problem
	if resp.Gid == 0 {
		return nil, nil
	}
	return
}

func (d *dao) AddNewGroup(ctx context.Context, req *grp.CreateGroupReq) (resp *grp.GroupInfo, err error) {
	resp, err = d.grpClient.CreateGroup(ctx, req)
	if err != nil {
		if ecode.Cause(err) == grp.GroupToAddNotExist {
			return
		}
		return resp, errors.Wrapf(err, "%+v", req)
	}
	return resp, nil
}

func (d *dao) AddNewMemberToGroup(ctx context.Context, uid int64, gid int64) error {
	_, err := d.grpClient.AddMember(ctx, &grp.AddMemberReq{
		Gid: gid,
		Uid: uid,
	})
	return err
}

func (d *dao) SearchGroupByName(ctx context.Context, name string) (groups *grp.AllGroups, err error) {
	groups, err = d.grpClient.GetAllGroupsLikeName(ctx, &grp.GroupsInfoByNameReq{
		Name: name,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "name : %s", name)
	}
	return groups, nil
}
