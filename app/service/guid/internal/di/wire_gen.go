// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"chat/app/service/guid/internal/dao"
	"chat/app/service/guid/internal/server/grpc"
	"chat/app/service/guid/internal/server/http"
	"chat/app/service/guid/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	db, cleanup, err := dao.NewDB()
	if err != nil {
		return nil, nil, err
	}
	daoDao, cleanup2, err := dao.New(db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serviceService, cleanup3, err := service.New(daoDao)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	engine, err := http.New(serviceService)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, err := grpc.New(serviceService)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup4, err := NewApp(serviceService, engine, server)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}