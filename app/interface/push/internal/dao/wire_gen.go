// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dao

// Injectors from wire.go:

func newTestDao() (*dao, func(), error) {
	rabbitService, cleanup, err := NewRabbitMqService()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := NewDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	daoDao, cleanup3, err := newDao(rabbitService, db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return daoDao, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
