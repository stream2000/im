package service

import (
	pb "chat/app/interface/push/api"
	"chat/app/interface/push/internal/dao"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.PushServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) PushUser(ctx context.Context, req *pb.PushUserReq) (*empty.Empty, error) {
	bmCtx := ctx.(*bm.Context)
	uid, ok := bmCtx.Get("uid")
	if !ok {
		return nil, ecode.ServerErr
	}
	req.SenderId = uid.(int64)
	err := s.dao.Send(req)
	resp := new(empty.Empty)
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
