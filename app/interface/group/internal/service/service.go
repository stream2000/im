package service

import (
	pb "chat/app/interface/group/api"
	"chat/app/interface/group/internal/dao"
	grp "chat/app/service/group/api"
	"context"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
	"google.golang.org/appengine/log"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.GroupServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) GetAllGroupsByUid(ctx context.Context, req *empty.Empty) (*pb.AllGroups, error) {
	bmCtx := ctx.(*bm.Context)
	uid, ok := bmCtx.Get("uid")
	if !ok {
		return nil, ecode.ServerErr
	}
	var groups *grp.AllGroups
	var err error
	groups, err = s.dao.AllGroupsByUserId(ctx, uid.(int64))
	if err != nil {
		return nil, ecode.Errorf(ecode.ServerErr, "error getting all groups")
	}
	if len(groups.Groups) == 0 {
		return nil, ecode.Errorf(ecode.NothingFound, "can't find any group")
	}
	resp := new(pb.AllGroups)
	for _, g := range groups.Groups {
		basicInfo := new(pb.GroupBasicInfo)
		basicInfo.Name = g.Name
		basicInfo.Gid = g.Gid
		basicInfo.Description = g.Description
		resp.Groups = append(resp.Groups, basicInfo)
	}
	return resp, nil
}

func (s *Service) CreateGroup(ctx context.Context, req *pb.CreateGroupReq) (*pb.GroupInfo, error) {
	bmCtx := ctx.(*bm.Context)
	uid, ok := bmCtx.Get("uid")

	if !ok {
		return nil, ecode.RequestErr
	}

	info, err := s.dao.AddNewGroup(ctx, &grp.CreateGroupReq{
		Uid:         uid.(int64),
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		log.Errorf(ctx, "error %+v creating group", err)
		return nil, ecode.Error(ecode.ServerErr, "error creating group")
	}
	resp := new(pb.GroupInfo)
	if info.Gid == 0 {
		log.Errorf(ctx, "error %+v creating group", err)
		return nil, ecode.Error(ecode.ServerErr, "error creating group: wrong value")
	}
	resp.Gid = info.Gid
	resp.Description = info.Description
	resp.Name = info.Name
	resp.Members = info.Members
	resp.MemberNumber = 1
	return resp, nil
}

func (s *Service) GetGroupInfo(ctx context.Context, req *pb.GroupInfoByIdReq) (*pb.GroupInfo, error) {
	info, err := s.dao.GerGroupDetailedInfo(ctx, req.Gid)
	if err != nil {
		return nil, ecode.Errorf(ecode.ServerErr, "error getting info for gid %d ", req.Gid)
	}
	if info.Gid == 0 {
		return nil, ecode.Errorf(ecode.NothingFound, "can't find group with id %d", req.Gid)
	}

	resp := new(pb.GroupInfo)

	resp.Name = info.Name
	resp.Description = info.Description
	resp.Gid = info.Gid
	resp.Members = info.Members
	return resp, nil
}

func (s *Service) GetAllGroups(ctx context.Context, req *pb.SearchGroupReq) (*pb.AllGroups, error) {
	// all groups
	var groups *grp.AllGroups
	var err error
	if req.Name == "" {
		groups, err = s.dao.AllGroups(ctx)
		if err != nil {
			return nil, ecode.Errorf(ecode.ServerErr, "error getting all groups")
		}
		if len(groups.Groups) == 0 {
			return nil, ecode.Errorf(ecode.NothingFound, "can't find any group")
		}
	} else {
		groups, err = s.dao.SearchGroupByName(ctx, req.Name)
		if err != nil {
			return nil, ecode.Errorf(ecode.ServerErr, "error searching groups")
		}
		if len(groups.Groups) == 0 {
			return nil, ecode.Errorf(ecode.NothingFound, "can't find any group with name %s", req.Name)
		}
	}
	resp := new(pb.AllGroups)
	for _, g := range groups.Groups {
		basicInfo := new(pb.GroupBasicInfo)
		basicInfo.Name = g.Name
		basicInfo.Gid = g.Gid
		basicInfo.Description = g.Description
		resp.Groups = append(resp.Groups, basicInfo)
	}
	return resp, nil
}

func (s *Service) AddMember(ctx context.Context, req *pb.AddMemberReq) (*empty.Empty, error) {
	uid, ok := ctx.(*bm.Context).Get("uid")
	resp := new(empty.Empty)
	if !ok {
		return nil, ecode.ServerErr
	}
	err := s.dao.AddNewMemberToGroup(ctx, uid.(int64), req.Gid)
	if err != nil {
		return resp, ecode.Errorf(ecode.ServerErr, "error add member: uid :%d gid %d ", uid.(int64), req.Gid)
	}
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
