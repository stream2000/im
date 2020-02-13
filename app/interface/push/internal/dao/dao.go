package dao

import (
	pb "chat/app/interface/push/api"
	"context"
	services "github.com/houseofcat/turbocookedrabbit/service"
	"time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"

	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRabbitMqService)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	Send(req *pb.PushUserReq) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
}

// dao dao.
type dao struct {
	db           *sql.DB
	cache        *fanout.Fanout
	rabbitSrv    *services.RabbitService
	demoExpire   int32
	exchangeName string
	routingKey   string
}

// New new a dao and return.
func New(srv *services.RabbitService, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(srv, db)
}

func newDao(srv *services.RabbitService, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire   xtime.Duration
		ExchangeName string
		RoutingKey   string
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:           db,
		cache:        fanout.New("cache"),
		rabbitSrv:    srv,
		exchangeName: cfg.ExchangeName,
		routingKey:   cfg.RoutingKey,
		demoExpire:   int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
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
