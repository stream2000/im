package service

import (
	"chat/app/common/tool/jwt"
	pb "chat/app/service/im_core/api"
	"chat/app/service/im_core/internal/dao"
	"chat/app/service/im_core/internal/model"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/prometheus/common/log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.ImCoreServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) Connect(ctx context.Context, req *pb.ConnectReq) (resp *pb.ConnectResp, err error) {
	resp = new(pb.ConnectResp)
	token := req.Jwt
	claims, err := jwt.ParseToken(token)

	if err != nil {
		log.Error("jwt.ParseToken(Token %s) error (%v)", token, err)
		return resp, ecode.Error(ecode.ServerErr, "jwt解析错误")
	}

	record := model.OnlineRecord{}
	record.DeviceId = claims.DeviceId
	record.DeviceType = claims.DeviceType
	record.Server = req.Server

	err = s.dao.AddUserKeyOnline(ctx, claims.Uid, claims.Id, record)
	resp.DeviceId = claims.DeviceId
	resp.Uid = claims.Uid
	resp.TokenId = claims.Id
	return resp, err
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

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
