package http

import (
	"chat/app/common/middleware/auth"
	pb "chat/app/interface/push/api"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// New new a bm server.
func New(s pb.PushServer) (engine *bm.Engine, err error) {
	var (
		cfg       bm.ServerConfig
		ct        paladin.TOML
		jwtConfig struct {
			JwtSecret string
		}
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&jwtConfig); err != nil {
		return
	}
	midMap := map[string]bm.HandlerFunc{
		"auth": auth.BearerAuth(jwtConfig.JwtSecret, 10),
	}
	engine = bm.DefaultServer(&cfg)
	pb.RegisterPushBMServer(engine, s, midMap)
	err = engine.Start()
	return
}
