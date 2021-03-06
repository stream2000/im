// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dao

// Injectors from wire.go:

func newTestDao() (*dao, func(), error) {
	consumer, cleanup, err := NewRabbitMqConsumer()
	if err != nil {
		return nil, nil, err
	}
	redis, cleanup2, err := NewRedis()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	memcache, cleanup3, err := NewMC()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db, cleanup4, err := NewDB()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	daoDao, cleanup5, err := newDao(consumer, redis, memcache, db)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return daoDao, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
