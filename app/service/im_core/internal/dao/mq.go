/*
@Time : 2020/2/14 13:11
@Author : Minus4
*/
package dao

import (
	"chat/app/service/im_core/internal/model"
	"context"
	"encoding/json"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	rabbitConsumer "github.com/houseofcat/turbocookedrabbit/consumer"
	rabbitModels "github.com/houseofcat/turbocookedrabbit/models"
	"github.com/houseofcat/turbocookedrabbit/pools"
	"time"
)

var channelPoolError, connectionPoolError <-chan error

func NewRabbitMqConsumer() (consumer *rabbitConsumer.Consumer, cf func(), err error) {
	var (
		Cfg struct {
			PoolConfig     *rabbitModels.PoolConfig
			ConsumerConfig *rabbitModels.ConsumerConfig
		}
	)
	if err = paladin.Get("rabbitmq.toml").UnmarshalTOML(&Cfg); err != nil {
		return
	}
	connectionPool, err := pools.NewConnectionPool(Cfg.PoolConfig, false)
	if err != nil {
		log.Error("Error Creating RabbitMq Connection Pool")
		return
	}
	channelPool, err := pools.NewChannelPool(Cfg.PoolConfig, connectionPool, false)
	if err != nil {
		log.Error("Error Creating RabbitMq Channel Pool")
		return
	}
	err = connectionPool.Initialize()
	if err != nil {
		log.Error("Error Initializing RabbitMq Connection Pool")
		return
	}
	err = channelPool.Initialize()
	if err != nil {
		log.Error("Error Initializing RabbitMq Channel Pool")
		return
	}
	consumer, err = rabbitConsumer.NewConsumerFromConfig(Cfg.ConsumerConfig, channelPool)
	if err != nil {
		log.Error("Error Creating RabbitMq Consumer")
		return
	}

	err = consumer.StartConsuming()
	if err != nil {
		log.Error("Error Initializing RabbitMq Consumer")
		return
	}

	cf = func() {
		_ = consumer.StopConsuming(true, true)
		channelPool.Shutdown()
	}
	channelPoolError = channelPool.Errors()
	connectionPoolError = connectionPool.Errors()
	return
}

func (d *dao) consume() {
	for {
		select {
		case err := <-channelPoolError:
			log.Error("%s: ChannelPool Error - %s", time.Now(), err)
		case err := <-connectionPoolError:
			log.Error("%s: ConnectionPool Error - %s", time.Now(), err)
		case err := <-d.consumer.Errors():
			log.Error("%s: Consumer Error - %s", time.Now(), err)
		case message := <-d.consumer.Messages():
			err := d.pool.Invoke(message)
			if err != nil {
				log.Error("error invoke pool : %v", err)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (d *dao) messageTask(arg interface{}) {
	msg := arg.(*rabbitModels.Message)
	var err error

	// clean up
	defer func() {
		if err != nil {
			err = msg.Nack(false)
			if err != nil {
				log.Error("dao.consume requeue error :%s", err)
			}
		} else {
			err = msg.Acknowledge()
			if err != nil {
				log.Error("%s: AckMessage Error - %s", time.Now(), err)
			}
		}
	}()

	var userMsg model.Message
	err = json.Unmarshal(msg.Body, &userMsg)
	if err != nil {
		log.Error("json.Unmarshal Error :(%v)", err)
		return
	}

	id, err := d.worker.Generate()
	if err != nil {
		log.Error("worker.generate error : (%v)",err)
		panic(err)
	}
	redisScore := d.worker.RedisScoreMapping(id)

	seq, err := d.NextTimelineSeq(keyC2CTimelineSeq(userMsg.SenderId, userMsg.ReceiverId))
	if err != nil {
		log.Error("dao.NextTimelineSeq error : (%v)",err)
		return
	}
	log.Info("receive msg : %+v with guid %d and seq %d", userMsg, id, seq)

	item := model.TimelineItemMap(userMsg)
	item.Guid = id
	item.ServerSeq = seq
	item.BusinessType = model.C2CMessage
	// store one item per timeline
	err = d.TimelineItemStore(keyC2CTimeline(item.SenderId, item.ReceiverId), redisScore, item)

	// #1 query redis whether the receiver is online
	servers,err := d.KeysByMid(context.Background(),userMsg.ReceiverId)
	if err != nil {
		log.Error("dao.KeysByMid(mid %d) error : (%v)",userMsg.ReceiverId,err)
		return
	}

	if len(servers) == 0 {
		log.Info("no online device of receiver id %d",userMsg.ReceiverId)
		// no online device
		return
	}
	// #2 rpc the comet to push the msg to receiver
	// NOTE here the worst part of my project
	for k,v := range servers{
		log.Info("push message to mid %d with  key %s in server %s",userMsg.ReceiverId,k,v.Server)
	}
	return
}
