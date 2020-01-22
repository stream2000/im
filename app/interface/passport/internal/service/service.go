package service

import (
	"chat/app/common/tool/basicAuth"
	"chat/app/common/tool/jwt"
	pb "chat/app/interface/passport/api"
	"chat/app/interface/passport/internal/dao"
	"context"
	"encoding/base64"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.PassportServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	email := req.Email
	passwordBytes, err := base64.StdEncoding.DecodeString(req.Password)
	if err != nil {
		return nil, ecode.Error(ecode.RequestErr, "密码编码错误")
	}
	password := string(passwordBytes)

	sum := basicAuth.EncryptAccount(email, password)
	uid, err := s.dao.Register(ctx, email, sum)

	if err != nil {
		return
	}

	resp = new(pb.RegisterResp)
	resp.Uid = uid

	return
}

func (s *Service) Login(ctx context.Context, req *google_protobuf1.Empty) (resp *pb.LoginResp, err error) {
	bmCtx := ctx.(*bm.Context)
	emailV, ok := bmCtx.Get("email")
	if !ok {
		panic(ecode.ServerErr)
	}
	email := emailV.(string)
	passwordV, ok := bmCtx.Get("password")
	if !ok {
		panic(ecode.ServerErr)
	}
	password := passwordV.(string)

	computedSum := basicAuth.EncryptAccount(email, password)
	authInfo, err := s.dao.AuthInfo(ctx, email)

	if err != nil {
		return
	}

	if computedSum != authInfo.Sum {
		return nil, ecode.Unauthorized
	}
	resp = &pb.LoginResp{}
	resp.Uid = authInfo.Uid
	resp.Token, _ = jwt.GenerateToken(resp.Uid)
	return
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// Close close the resource.
func (s *Service) Close() {
}
