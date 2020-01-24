// Code generated by kratos tool genmc. DO NOT EDIT.

/*
  Package dao is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=keyBasic -type=get
		CacheBasicInfo(c context.Context, id int64) (*pb.BasicInfo, error)
		// mc: -key=keyBasic -expire=d.demoExpire
		AddCacheBasicInfo(c context.Context, id int64, art *pb.BasicInfo) (err error)
		// mc: -key=keyBasic
		DeleteBasicInfoCache(c context.Context, id int64) (err error)
	}
*/

package dao

import (
	"context"
	"fmt"

	pb "chat/app/interface/account/api"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/log"
)

var _ _mc

// CacheBasicInfo get data from mc
func (d *dao) CacheBasicInfo(c context.Context, id int64) (res *pb.BasicInfo, err error) {
	key := keyBasic(id)
	res = &pb.BasicInfo{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheBasicInfo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheBasicInfo Set data to mc
func (d *dao) AddCacheBasicInfo(c context.Context, id int64, val *pb.BasicInfo) (err error) {
	if val == nil {
		return
	}
	key := keyBasic(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheBasicInfo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DeleteBasicInfoCache delete data from mc
func (d *dao) DeleteBasicInfoCache(c context.Context, id int64) (err error) {
	key := keyBasic(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DeleteBasicInfoCache", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
