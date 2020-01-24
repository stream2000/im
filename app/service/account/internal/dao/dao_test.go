package dao

import (
	pb "chat/app/service/account/api"
	"context"
	"flag"
	"github.com/bilibili/kratos/pkg/ecode"
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
func TestDao_Account(t *testing.T) {
	res, err := d.Account(ctx, 1)
	if err != nil {
		panic(err)
	}
	if res == nil {
		panic(ecode.NothingFound)
	}

	log.Infoc(ctx, "res: %v", res)

	newEmail := "4600332@gmail.com"
	newRegisterReq := &pb.RegisterReq{
		Email:    newEmail,
		Password: "1234",
	}

	uid, err := d.AddAccount(ctx, newRegisterReq)

	if err != nil {
		panic(err)
	} else {
		log.Infoc(ctx, "UID: %s", uid)
	}

	addedAcc, err := d.Account(ctx, 1)

	if err != nil {
		panic(err)
	}
	if addedAcc == nil {
		log.Infoc(ctx, "Nothing found with email %s", newEmail)
		panic(ecode.NothingFound)
	}
}
