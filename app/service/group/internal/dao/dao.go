package dao

import (
	pb "chat/app/service/group/api"
	"context"
	"time"

	"chat/app/service/group/internal/model"
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
	GetAllGroupsByUserId(ctx context.Context, uid int64) (groups []*model.Group, err error)
	GetAllGroups(ctx context.Context) (groups []*model.Group, err error)
	CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (info *pb.GroupInfo, err error)
	// bts:  -nullcache=&model.Group{Id:0} -check_null_code=$!=nil&&$.Id==0 -sync=true
	Group(ctx context.Context, gid int64) (*model.Group, error)
	GetAllGroupsByName(ctx context.Context, name string) (groups []*model.Group, err error)
	AddMember(ctx context.Context, uid int64, gid int64) error
}

// dao dao.
type dao struct {
	db         *sql.DB
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
	d = &dao{
		db:         db,
		redis:      r,
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
