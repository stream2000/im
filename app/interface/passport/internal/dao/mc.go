package dao

import (
	"chat/app/interface/passport/internal/model"
	"context"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

//go:generate kratos tool genmc
type _mc interface {
	// mc: -key=keyAcInfo -type=get
	CacheAuthInfo(c context.Context, email string) (resp *model.AuthInfo, err error)
	// mc: -key=keyAcInfo -expire=d.demoExpire
	AddCacheAuthInfo(c context.Context, email string, info *model.AuthInfo) (err error)
	// mc: -key=keyAcInfo
	DeleteCacheAuthInfo(c context.Context, email string) (err error)
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

func keyAcInfo(email string) string {
	return "acinfo_" + email
}
