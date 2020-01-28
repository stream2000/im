// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		Register(ctx context.Context, email, password string) (uid int64, err error)
		// bts: -nullcache=&model.AuthInfo{Uid:0} -check_null_code=$!=nil&&$.Uid==0
		AuthInfo(ctx context.Context, email string) (resp *model.AuthInfo, err error)
	}
*/

package dao

import (
	"context"

	"chat/app/interface/passport/internal/model"
	"github.com/bilibili/kratos/pkg/cache"
)

// AuthInfo get data from cache if miss will call source method, then add to cache.
func (d *dao) AuthInfo(c context.Context, email string) (res *model.AuthInfo, err error) {
	addCache := true
	res, err = d.CacheAuthInfo(c, email)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.Uid == 0 {
			res = nil
		}
	}()
	if res != nil {
		cache.MetricHits.Inc("bts:AuthInfo")
		return
	}
	cache.MetricMisses.Inc("bts:AuthInfo")
	res, err = d.RawAuthInfo(c, email)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.AuthInfo{Uid: 0}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheAuthInfo(c, email, miss)
	})
	return
}
