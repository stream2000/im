// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"chat/app/interface/push/internal/dao"
	"chat/app/interface/push/internal/server/http"
	"chat/app/interface/push/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	rabbitService, cleanup, err := dao.NewRabbitMqService()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := dao.NewDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	daoDao, cleanup3, err := dao.New(rabbitService, db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceService, cleanup4, err := service.New(daoDao)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	engine, err := http.New(serviceService)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup5, err := NewApp(serviceService, engine)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
