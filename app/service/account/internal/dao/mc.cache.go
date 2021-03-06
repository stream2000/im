// Code generated by kratos tool genmc. DO NOT EDIT.

/*
  Package dao is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=keyAcc -type=get
		CacheAccount(c context.Context, id int64) (*model.Account, error)
		// mc: -key=keyAcc -expire=d.demoExpire
		AddCacheAccount(c context.Context, id int64, acc *model.Account) (err error)
		// mc: -key=keyAcc
		DeleteCacheAccount(c context.Context, id int64) (err error)
	}
*/

package dao

import (
	"context"
	"fmt"

	"chat/app/service/account/internal/model"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/log"
)

var _ _mc

// CacheAccount get data from mc
func (d *dao) CacheAccount(c context.Context, id int64) (res *model.Account, err error) {
	key := keyAcc(id)
	res = &model.Account{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheAccount", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheAccount Set data to mc
func (d *dao) AddCacheAccount(c context.Context, id int64, val *model.Account) (err error) {
	if val == nil {
		return
	}
	key := keyAcc(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheAccount", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DeleteCacheAccount delete data from mc
func (d *dao) DeleteCacheAccount(c context.Context, id int64) (err error) {
	key := keyAcc(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DeleteCacheAccount", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
