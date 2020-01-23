package dao

import (
	"context"
	"fmt"

	"chat/app/service/group/internal/model"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

//go:generate kratos tool genmc
type _mc interface {
	// mc: -key=keyGroup -type=get
	CacheGroup(c context.Context, id int64) (*model.Group, error)
	// mc: -key=keyGroup -expire=d.demoExpire
	AddCacheGroup(c context.Context, id int64, art *model.Group) (err error)
	// mc: -key=keyGroup
	DeleteGroupCache(c context.Context, id int64) (err error)
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

func keyGroup(id int64) string {
	return fmt.Sprintf("g%d", id)
}
