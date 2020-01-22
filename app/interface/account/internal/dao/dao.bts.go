// Code generated by kratos tool genbts. DO NOT EDIT.

/*
  Package dao is a generated cache proxy package.
  It is generated from:
  type Dao interface {
		Close()
		Ping(ctx context.Context) (err error)
		// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
		Article(c context.Context, id int64) (*model.Article, error)
	}
*/

package dao

import (
	"context"

	"chat/app/interface/account/internal/model"
	"github.com/bilibili/kratos/pkg/cache"
)

// Article get data from cache if miss will call source method, then add to cache.
func (d *dao) Article(c context.Context, id int64) (res *model.Article, err error) {
	addCache := true
	res, err = d.CacheArticle(c, id)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res != nil && res.ID == -1 {
			res = nil
		}
	}()
	if res != nil {
		cache.MetricHits.Inc("bts:Article")
		return
	}
	cache.MetricMisses.Inc("bts:Article")
	res, err = d.RawArticle(c, id)
	if err != nil {
		return
	}
	miss := res
	if miss == nil {
		miss = &model.Article{ID: -1}
	}
	if !addCache {
		return
	}
	d.cache.Do(c, func(c context.Context) {
		d.AddCacheArticle(c, id, miss)
	})
	return
}
