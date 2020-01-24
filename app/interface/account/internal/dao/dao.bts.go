// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		// bts: -check_null_code=$!=nil&&$.Uid==0
		BasicInfo(ctx context.Context, id int64)(*pb.BasicInfo,error)
	}
*/

package dao

import (
	"context"

	pb "chat/app/interface/account/api"
	"github.com/bilibili/kratos/pkg/cache"
)

// BasicInfo get data from cache if miss will call source method, then add to cache.
func (d *dao) BasicInfo(c context.Context, id int64) (res *pb.BasicInfo, err error) {
	addCache := true
	res, err = d.CacheBasicInfo(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	if res != nil {
		cache.MetricHits.Inc("bts:BasicInfo")
		return
	}
	cache.MetricMisses.Inc("bts:BasicInfo")
	res, err = d.RawBasicInfo(c, id)
	if err != nil {
		return
	}
	miss := res
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheBasicInfo(c, id, miss)
	})
	return
}
