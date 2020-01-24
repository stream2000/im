package service

import (
	pb "chat/app/interface/account/api"
	"chat/app/interface/account/internal/dao"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.AccountServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) GetBasicInfo(ctx context.Context, req *empty.Empty) (*pb.BasicInfo, error) {
	c := ctx.(*bm.Context)
	uidValue, ok := c.Get("uid")
	uid := uidValue.(int64)
	if !ok {
		log.Error("can;t load uid after jwt authentication")
		panic("Auth Error")
	}
	info, err := s.dao.BasicInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, ecode.NothingFound
	}

	resp := new(pb.BasicInfo)

	resp.ProfilePicUrl = info.ProfilePicUrl
	resp.Uid = info.Uid
	resp.Email = info.Email
	resp.Sign = info.Sign
	resp.Nickname = info.Nickname
	return resp, nil
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
