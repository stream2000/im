package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/google/wire"
	"time"
)

var Provider = wire.NewSet(New, NewDB)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	GetGUID(ctx context.Context) (int64, error)
}

// dao dao.
type dao struct {
	db       *sql.DB
	ch       chan int64 // id缓冲池
	min, max int64      // id段最小值，最大值
	cfg      *daoCfg
}
type daoCfg struct {
	DemoExpire xtime.Duration
	BufferLen  int
	Step       int64
}

// New new a dao and return.
func New(db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(db)
}

func newDao(db *sql.DB) (d *dao, cf func(), err error) {
	var ac daoCfg
	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		return
	}
	log.Info(" cfg : %+v", ac)
	d = &dao{
		db:  db,
		cfg: &ac,
		ch:  make(chan int64, ac.BufferLen),
	}
	cf = d.Close
	go d.productId()
	return
}

func (d *dao) GetGUID(ctx context.Context) (int64, error) {
	select {
	case <-time.After(1 * time.Second):
		return 0, ecode.Errorf(ecode.ServerErr, "timeout fetch id form db")
	case uid := <-d.ch:
		return uid, nil
	}
}

func (d *dao) reLoad() {
	var err error
	for {
		err = d.fetchGuid(context.Background())
		if err == nil {
			return
		}
		log.Error("dao.reload : Fatal error when fetching id from mysql")
		time.Sleep(time.Second)
	}
}

func (d *dao) productId() {
	d.reLoad()
	for {
		if d.min >= d.max {
			log.Info("here min : %d max %d ", d.min, d.max)
			d.reLoad()
		}
		d.ch <- d.min
		d.min++
	}
}
func (d *dao) fetchGuid(ctx context.Context) (err error) {
	var step = d.cfg.Step
	var maxId int64
	tx, err := d.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow("select max_id from generator_table")
	err = row.Scan(&maxId)
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE generator_table SET max_id = ?", maxId+step)

	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	d.min = maxId
	d.max = maxId + step
	return
}

// Close close the resource.
func (d *dao) Close() {

}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
