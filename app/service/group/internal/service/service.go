package service

import (
	pb "chat/app/service/group/api"
	"chat/app/service/group/internal/dao"
	"chat/app/service/group/internal/model"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.GroupServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) GetAllGroupsLikeName(ctx context.Context, req *pb.GroupsInfoByNameReq) (allGroups *pb.AllGroups, err error) {
	allGroups = new(pb.AllGroups)
	groups, err := s.dao.GetAllGroupsByName(ctx, req.Name)
	if err != nil {
		return allGroups, err
	}
	if groups == nil {
		return allGroups, nil
	}
	for _, g := range groups {
		basicInfo := new(pb.GroupBasicInfo)
		basicInfo.Gid = g.Id
		basicInfo.Name = g.Name
		basicInfo.Description = g.Description
		allGroups.Groups = append(allGroups.Groups, basicInfo)
	}
	return allGroups, nil
}

func (s *Service) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.GroupInfo, error) {
	info := new(pb.GroupInfo)
	info, err := s.dao.CreateGroup(ctx, req)
	return info, err
}

func (s *Service) GetGroupInfo(ctx context.Context, req *pb.GroupInfoByIdReq) (resp *pb.GroupInfo, err error) {
	resp = new(pb.GroupInfo)
	g, err := s.dao.Group(ctx, req.Gid)
	if err != nil {
		return resp, ecode.Errorf(ecode.ServerErr, "error when get group info by id %s", err.Error())
	}
	if g == nil {
		return resp, nil
	}
	resp.Gid = g.Id
	resp.Name = g.Name
	resp.Description = g.Description
	resp.Members = g.Members
	return
}

func (s *Service) GetAllGroups(ctx context.Context, req *pb.GetAllGroupsReq) (resp *pb.AllGroups, err error) {
	resp = new(pb.AllGroups)
	var groups []*model.Group
	if req.Uid == 0 {
		groups, err = s.dao.GetAllGroups(ctx)
	} else {
		groups, err = s.dao.GetAllGroupsByUserId(ctx, req.Uid)
	}
	if err != nil {
		return resp, ecode.Errorf(ecode.ServerErr, "error when get group info by id %s", err.Error())
	}
	if len(groups) == 0 {
		return resp, nil
	}
	for _, g := range groups {
		basicInfo := new(pb.GroupBasicInfo)
		basicInfo.Gid = g.Id
		basicInfo.Name = g.Name
		basicInfo.Description = g.Description
		resp.Groups = append(resp.Groups, basicInfo)
	}
	return
}

func (s *Service) AddMember(ctx context.Context, req *pb.AddMemberReq) (*empty.Empty, error) {
	resp := new(empty.Empty)
	err := s.dao.AddMember(ctx, req.Uid, req.Gid)
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

// Close close the resource.
func (s *Service) Close() {

}
