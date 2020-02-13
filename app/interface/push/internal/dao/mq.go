/*
@Time : 2020/2/13 14:27
@Author : Minus4
*/
package dao

import (
	pb "chat/app/interface/push/api"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	rabbitModels "github.com/houseofcat/turbocookedrabbit/models"
	rabbitSrv "github.com/houseofcat/turbocookedrabbit/service"
)

func NewRabbitMqService() (srv *rabbitSrv.RabbitService, cf func(), err error) {
	var (
		cfg rabbitModels.RabbitSeasoning
	)
	if err = paladin.Get("rabbitmq.toml").UnmarshalTOML(&cfg); err != nil {
		fmt.Println("error")
		return
	}
	srv, err = rabbitSrv.NewRabbitService(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	srv.StartService(true)
	cf = srv.StopService
	return
}

func (d *dao) Send(req *pb.PushUserReq) (err error) {
	err = d.rabbitSrv.Publish(req, d.exchangeName, d.routingKey, false, "")
	return err
}
