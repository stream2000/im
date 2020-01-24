package dao

import (
	pb "chat/app/interface/account/api"
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

//go:generate kratos tool genmc
type _mc interface {
	// mc: -key=keyBasic -type=get
	CacheBasicInfo(c context.Context, id int64) (*pb.BasicInfo, error)
	// mc: -key=keyBasic -expire=d.demoExpire
	AddCacheBasicInfo(c context.Context, id int64, art *pb.BasicInfo) (err error)
	// mc: -key=keyBasic
	DeleteBasicInfoCache(c context.Context, id int64) (err error)
}

func NewMC() (mc *memcache.Memcache, cf func(), err error) {
	var (
		cfg memcache.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("memcache.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	mc = memcache.New(&cfg)
	cf = func() { mc.Close() }
	return
}

func (d *dao) PingMC(ctx context.Context) (err error) {
	if err = d.mc.Set(ctx, &memcache.Item{Key: "ping", Value: []byte("pong"), Expiration: 0}); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func keyBasic(id int64) string {
	return fmt.Sprintf("art_%d", id)
}
