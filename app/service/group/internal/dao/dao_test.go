package dao

import (
	pb "chat/app/service/group/api"
	"context"
	"flag"
	"github.com/bilibili/kratos/pkg/log"
	"os"
	"testing"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/testing/lich"
)

var d *dao
var ctx = context.Background()

func TestMain(m *testing.M) {
	flag.Set("conf", "../../test")
	flag.Set("f", "../../test/docker-compose.yaml")
	flag.Parse()
	disableLich := os.Getenv("DISABLE_LICH") != ""
	if !disableLich {
		if err := lich.Setup(); err != nil {
			panic(err)
		}
	}
	var err error
	if err = paladin.Init(); err != nil {
		panic(err)
	}
	var cf func()
	if d, cf, err = newTestDao(); err != nil {
		panic(err)
	}
	ret := m.Run()
	cf()
	if !disableLich {
		_ = lich.Teardown()
	}
	os.Exit(ret)
}

func TestAll(t *testing.T) {
	info, err := d.CreateGroup(ctx, &pb.CreateGroupReq{
		Uid:         1,
		Name:        "new",
		Description: "test",
	})

	if err != nil {
		panic(err)
	}
	log.Info("%+v", *info)
	//g, err := d.Group(ctx, 2)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//log.Infoc(ctx, "%+v", g)
	//
	//err = d.AddMember(ctx, 100, 1)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//groups, err := d.GetAllGroupsByUserId(ctx, 1)
	//
	//if err != nil {
	//	panic(groups)
	//}
	//for _, g := range groups {
	//	log.Info("groups : %+v", *g)
	//}
	//
	//groups, err = d.GetAllGroupsByName(ctx, "test")
	//
	//if err != nil {
	//	panic(groups)
	//}
	//for _, g := range groups {
	//	log.Info("groups : %+v", *g)
	//}
}
