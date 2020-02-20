package dao

import (
	"chat/app/service/im_core/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

const (
	_prefixMidServer      = "mid_%d" // mid -> key:server
	_prefixKeyServer      = "key_%s" // key -> server
	_prefixServerOnline   = "ol_%s"  // server -> online
	_prefixC2CTimelineSeq = "c2ctlSeq_%d_to_%d"
	_prefixC2CTimeline    = "c2ctl%d_to_%d"
	_prefixC2GTimelineSeq = "c2gtlSeq_%d_to_%d"
	_prefixC2GTimeline    = "c2gtl%d_to_%d"
)

func NewRedis() (r *redis.Redis, cf func(), err error) {
	var (
		cfg redis.Config
		ct  paladin.Map
	)
	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Client").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewRedis(&cfg)
	cf = func() { r.Close() }
	return
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	if _, err = d.redis.Do(ctx, "SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}

func keyMidServer(mid int64) string {
	return fmt.Sprintf(_prefixMidServer, mid)
}

func keyKeyServer(key string) string {
	return fmt.Sprintf(_prefixKeyServer, key)
}

func keyServerOnline(key string) string {
	return fmt.Sprintf(_prefixServerOnline, key)
}

func keyC2CTimelineSeq(senderId, receiverId int64) string {
	return fmt.Sprintf(_prefixC2CTimelineSeq, senderId, receiverId)
}

func keyC2CTimeline(senderId, receiverId int64) string {
	return fmt.Sprintf(_prefixC2CTimeline, senderId, receiverId)
}
func keyC2GTimeline(senderId, receiverId int64) string {
	return fmt.Sprintf(_prefixC2GTimelineSeq, senderId, receiverId)
}

// bytesMap redis helper function, migrated from redis.StringMap
func bytesMap(result interface{}, err error) (map[string][]byte, error) {
	values, err := redis.Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: bytesMap expects even number of values result")
	}
	m := make(map[string][]byte, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, okKey := values[i].([]byte)
		value, okValue := values[i+1].([]byte)
		if !okKey || !okValue {
			return nil, errors.New("redigo: StringMap key not a bulk string value")
		}
		m[string(key)] = value
	}
	return m, nil
}

// key represents jwt uuid
func (d *dao) AddUserKeyOnline(ctx context.Context, mid int64, key string, record model.OnlineRecord) (err error) {
	var server = record.Server
	data, err := json.Marshal(record)
	if err != nil {
		log.Error("json.Marshal(Record %v) error(%v)", record, err)
		return
	}
	conn := d.redis.Conn(ctx)
	defer conn.Close()
	var n = 2
	if mid > 0 {
		if err = conn.Send("HSET", keyMidServer(mid), key, data); err != nil {
			return
		}
		if err = conn.Send("EXPIRE", keyMidServer(mid), d.demoExpire); err != nil {
			log.Error("conn.Send(EXPIRE %d,%s,%s) error(%v)", mid, key, server, err)
			return
		}
		n += 2
	}
	if err = conn.Send("SET", keyKeyServer(key), server); err != nil {
		log.Error("conn.Send(HSET %d,%s,%s) error(%v)", mid, server, key, err)
		return
	}
	if err = conn.Send("EXPIRE", keyKeyServer(key), d.demoExpire); err != nil {
		log.Error("conn.Send(EXPIRE %d,%s,%s) error(%v)", mid, key, server, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < n; i++ {
		if _, err = conn.Receive(); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

// DelUserKeyOnline del a mapping.
func (d *dao) DelUserKeyOnline(c context.Context, mid int64, key, server string) (has bool, err error) {
	conn := d.redis.Conn(c)
	defer conn.Close()
	n := 1
	if mid > 0 {
		if err = conn.Send("HDEL", keyMidServer(mid), key); err != nil {
			log.Error("conn.Send(HDEL %d,%s,%s) error(%v)", mid, key, server, err)
			return
		}
		n++
	}
	if err = conn.Send("DEL", keyKeyServer(key)); err != nil {
		log.Error("conn.Send(HDEL %d,%s,%s) error(%v)", mid, key, server, err)
		return
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}
	for i := 0; i < n; i++ {
		if has, err = redis.Bool(conn.Receive()); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
	}
	return
}

func (d *dao) KeysByMids(c context.Context, mids []int64) (ress map[string]string, olMids []int64, err error) {
	conn := d.redis.Conn(c)
	defer conn.Close()
	ress = make(map[string]string)
	// a redis pipeline operation
	for _, mid := range mids {
		// get the key-server map
		if err = conn.Send("HGETALL", keyMidServer(mid)); err != nil {
			log.Error("conn.Do(HGETALL %d) error(%v)", mid, err)
			return
		}
	}
	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}

	for idx := 0; idx < len(mids); idx++ {
		var (
			res map[string][]byte
		)

		if res, err = bytesMap(conn.Receive()); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
		if len(res) > 0 {
			olMids = append(olMids, mids[idx])
		}
		for k, v := range res {
			var reply model.OnlineRecord
			err = json.Unmarshal(v, &reply)
			if err != nil {
				log.Error("json.Unmarshal() error(%v)", err)
				return
			}
			ress[k] = reply.Server
		}
	}
	return
}

// use to kick out users with the same mid and device type
func (d *dao) KeysByMid(c context.Context, mid int64) (res map[string]model.OnlineRecord, err error) {
	conn := d.redis.Conn(c)
	defer conn.Close()
	rpl, err := bytesMap(conn.Do("HGETALL", keyMidServer(mid)))
	if err != nil {
		log.Error("conn.DO() error(%v)", err)
		return
	}
	if len(rpl) == 0 {
		return
	}
	res = make(map[string]model.OnlineRecord)
	for k, v := range rpl {
		var reply model.OnlineRecord
		err = json.Unmarshal(v, &reply)
		if err != nil {
			log.Error("json.Unmarshal() error(%v)", err)
			return
		}
		res[k] = reply
	}
	return
}

func (d *dao) OnlineList(c context.Context, mids []int64) (olMids []int64, err error) {
	conn := d.redis.Conn(c)
	defer conn.Close()
	for _, mid := range mids {
		if err = conn.Send("EXISTS", keyMidServer(mid)); err != nil {
			log.Error("conn.Do(EXISTS %d) error(%v)", mid, err)
			return
		}
	}

	if err = conn.Flush(); err != nil {
		log.Error("conn.Flush() error(%v)", err)
		return
	}

	for idx := 0; idx < len(mids); idx++ {

		var exists bool
		if exists, err = redis.Bool(conn.Receive()); err != nil {
			log.Error("conn.Receive() error(%v)", err)
			return
		}
		if exists {
			olMids = append(olMids, mids[idx])
		}
	}

	return
}

func (d *dao) NextTimelineSeq(timelineKey string) (int64, error) {
	redisConn := d.redis.Conn(context.Background())
	defer redisConn.Close()
	seq, err := redis.Int64(redisConn.Do("INCR", timelineKey))
	if err != nil {
		return 0, fmt.Errorf("err : %+v, redis incr", err)
	}
	return seq, nil
}

func (d *dao) TimelineItemStore(key string, score int64, item model.TimelineItem) (err error) {
	conn := d.redis.Conn(context.Background())
	defer conn.Close()
	data, err := json.Marshal(item)
	if err != nil {
		log.Error("json.Unmarshal Error :(%v)", err)
		return
	}
	_, err = conn.Do("zadd", key, score, data)
	if err != nil {
		log.Error("conn.Do(zadd key item (%+v)) error(%v)", key, item, err)
	}
	return err
}
