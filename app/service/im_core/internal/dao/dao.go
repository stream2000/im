package dao

import (
	"chat/app/common/tool/jwt"
	"chat/app/common/tool/snowflake"
	"chat/app/service/im_core/internal/model"
	"context"
	rabbitConsumer "github.com/houseofcat/turbocookedrabbit/consumer"
	"time"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/google/wire"
	"github.com/panjf2000/ants/v2"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC, NewRabbitMqConsumer)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	AddUserKeyOnline(ctx context.Context, mid int64, key string, record model.OnlineRecord) error
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	consumer   *rabbitConsumer.Consumer
	cache      *fanout.Fanout
	demoExpire int32
	pool       *ants.PoolWithFunc
	worker     *snowflake.Worker
}

// New new a dao and return.
func New(consumer *rabbitConsumer.Consumer, r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(consumer, r, mc, db)
}

func newDao(consumer *rabbitConsumer.Consumer, r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
		JwtSecret  string
	}
	d = new(dao)
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	// TODO get worker id from some kind of config file, for example, file , environment variable
	worker, err := snowflake.NewWorker(1)
	d = &dao{
		db:         db,
		redis:      r,
		mc:         mc,
		consumer:   consumer,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
		worker:     worker,
	}
	p, _ := ants.NewPoolWithFunc(50000, d.messageTask)
	d.pool = p
	jwt.Init(cfg.JwtSecret)
	cf = func() {
		p.Release()
		d.Close()
	}
	go d.consume()
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
