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
		appCfg   paladin.TOML
		authConf struct {
			Salt1     string `dsn:"salt1"`
			Salt2     string `dsn:"salt2"`
			JwtSecret string `dsn:"jwtSecret"`
		}
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
	if err = appCfg.Get("auth").UnmarshalTOML(&authConf); err != nil {
		return
	}
	basicAuth.Setup(authConf.Salt1, authConf.Salt2)
	jwt.Setup(authConf.JwtSecret)
	midMap := map[string]bm.HandlerFunc{
		"auth":  auth.BearerAuth(authConf.JwtSecret, 10),
		"basic": auth.BasicFilter,
	}
	engine = bm.DefaultServer(&cfg)
	pb.RegisterPassportBMServer(engine, s, midMap)
	err = engine.Start()
	return
}
