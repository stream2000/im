// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		// bts:  -nullcache=&model.Account{Email:"invalid"} -check_null_code=$!=nil&&$.Email=="invalid" -sync=true
		Account(c context.Context, email string) (*model.Account, error)
		AddAccount(c context.Context,req *pb.RegisterReq)(*pb.RegisterResp,error)
	}
*/

package dao

import (
	"context"

	"chat/app/service/account/internal/model"
	"github.com/bilibili/kratos/pkg/cache"
)

// Account get data from cache if miss will call source method, then add to cache.
func (d *dao) Account(c context.Context, email string) (res *model.Account, err error) {
	addCache := true
	res, err = d.CacheAccount(c, email)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.Email == "invalid" {
			res = nil
		}
	}()
	if res != nil {
		cache.MetricHits.Inc("bts:Account")
		return
	}
	cache.MetricMisses.Inc("bts:Account")
	res, err = d.RawAccount(c, email)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.Account{Email: "invalid"}
	}
	if !addCache {
		return
	}
	d.AddCacheAccount(c, email, miss)
	return
}