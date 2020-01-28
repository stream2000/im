package service

import (
	pb "chat/app/service/account/api"
	"chat/app/service/account/internal/dao"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/google/wire"
)

// The service struct is expected to implements the pb.AccountServer
var Provider = wire.NewSet(New, wire.Bind(new(pb.AccountServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) GetBasicInfo(ctx context.Context, req *pb.BasicInfoReq) (resp *pb.BasicInfo, err error) {
	resp = new(pb.BasicInfo)

	acc, err := s.dao.Account(ctx, req.Uid)
	if err != nil {
		return
	}
	if acc == nil {
		return
	}
	resp.Sign = acc.Sign
	resp.Uid = acc.UID
	resp.Email = acc.Email
	resp.ProfilePicUrl = acc.ProfilePicUrl
	resp.Nickname = acc.NickName
	return
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (rsp *pb.RegisterResp, err error) {
	rsp = &pb.RegisterResp{}
	rsp, err = s.dao.AddAccount(ctx, req)
	return rsp, err
}

func (s *Service) GetAuthInfo(ctx context.Context, req *pb.AuthReq) (resp *pb.AuthResp, err error) {
	resp = new(pb.AuthResp)

	acc, err := s.dao.GetAccountByEmail(ctx, req.Email)

	if err != nil {
		return
	}
	if acc == nil {
		return
	}

	resp.Password = acc.Password
	resp.Uid = acc.UID
	return
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	// so all config can read from the file
	err = paladin.Watch("application.toml", s.ac)
	return
}

// Close close the resource.
func (s *Service) Close() {
}
