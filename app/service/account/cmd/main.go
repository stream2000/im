package main

import (
	"chat/app/service/account/api"
	"context"
	"flag"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chat/app/service/account/internal/di"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("account start")

	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	var (
		cfg warden.ServerConfig
		ct  paladin.TOML
		ac  struct {
			Discovery *discovery.Config
		}
		cancel context.CancelFunc
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	if err := paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	dis := discovery.New(ac.Discovery)
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		Hostname: env.Hostname,
		AppID:    api.AppID,
		Addrs: []string{
			"grpc://" + cfg.Addr,
		},
	}
	cancel, err = dis.Register(context.Background(), ins)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("account exit")
			if cancel != nil {
				cancel()
			}
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
