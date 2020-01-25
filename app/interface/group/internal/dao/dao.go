package dao

import (
	grp "chat/app/service/group/api"
	"context"
	"github.com/prometheus/common/log"
	"time"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	GerGroupDetailedInfo(ctx context.Context, gid int64) (resp *grp.GroupInfo, err error)
	AllGroupsByUserId(ctx context.Context, uid int64) (resp *grp.AllGroups, err error)
	AllGroups(ctx context.Context) (resp *grp.AllGroups, err error)
	AddNewGroup(ctx context.Context, req *grp.CreateGroupReq) (resp *grp.GroupInfo, err error)
	AddNewMemberToGroup(ctx context.Context, uid int64, gid int64) error
	SearchGroupByName(ctx context.Context, name string) (groups *grp.AllGroups, err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
}

// dao dao.
type dao struct {
	db         *sql.DB
	grpClient  grp.GroupClient
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}

func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}

	grpClient, err := NewWardenClient()
	if err != nil {
		log.Error("Error dialing grpc server : ", err)
		panic(err)
	}
	d = &dao{
		db:         db,
		redis:      r,
		grpClient:  grpClient,
		mc:         mc,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}
