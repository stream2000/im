package server

import (
	"chat/app/common/middleware/auth"
	pb "chat/app/interface/account/api"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// New new a bm server.
func New(s pb.AccountBMServer) (engine *bm.Engine, err error) {
	var (
		cfg    bm.ServerConfig
		ct     paladin.TOML
		appCfg paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	if err = paladin.Get("application.toml").Unmarshal(&appCfg); err != nil {
		return
	}
	jwtSecret, err := appCfg.Get("jwtSecret").String()
	if err != nil {
		return
	}
	midMap := map[string]bm.HandlerFunc{
		"auth": auth.BearerAuth(jwtSecret, 10),
	}
	engine = bm.DefaultServer(&cfg)
	pb.RegisterAccountBMServer(engine, s, midMap)
	err = engine.Start()
	return
}
