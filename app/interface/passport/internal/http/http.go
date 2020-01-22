package http

import (
	"chat/app/common/middleware/auth"
	"chat/app/common/tool/basicAuth"
	"chat/app/common/tool/jwt"
	pb "chat/app/interface/passport/api"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// New new a bm http.
func New(s pb.PassportServer) (engine *bm.Engine, err error) {
	var (
		cfg      bm.ServerConfig
		ct       paladin.TOML
		authConf struct {
			jwtSecret string
			salt1     string
			salt2     string
		}
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&authConf); err != nil {
		return
	}

	basicAuth.Setup(authConf.salt1, authConf.salt2)
	jwt.Setup(authConf.jwtSecret)
	midMap := map[string]bm.HandlerFunc{
		"auth":  auth.BearerAuth(authConf.jwtSecret, 10),
		"basic": auth.BasicFilter,
	}
	engine = bm.DefaultServer(&cfg)
	pb.RegisterPassportBMServer(engine, s, midMap)
	err = engine.Start()
	return
}
