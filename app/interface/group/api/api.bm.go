// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

package api

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathGroupCreateGroup = "/group/create"
var PathGroupGetGroupInfo = "/group/info"
var PathGroupGetAllGroups = "/group/all"
var PathGroupGetAllGroupsByUid = "/group/all/user"
var PathGroupAddMember = "/group/addMember"

// GroupBMServer is the server API for Group service.
type GroupBMServer interface {
	// `midware:"auth"`
	CreateGroup(ctx context.Context, req *CreateGroupReq) (resp *GroupInfo, err error)

	GetGroupInfo(ctx context.Context, req *GroupInfoByIdReq) (resp *GroupInfo, err error)

	GetAllGroups(ctx context.Context, req *SearchGroupReq) (resp *AllGroups, err error)

	// `midware:"auth"`
	GetAllGroupsByUid(ctx context.Context, req *google_protobuf1.Empty) (resp *AllGroups, err error)

	// `midware:"auth"`
	AddMember(ctx context.Context, req *AddMemberReq) (resp *google_protobuf1.Empty, err error)
}

var GroupSvc GroupBMServer

func groupCreateGroup(c *bm.Context) {
	p := new(CreateGroupReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GroupSvc.CreateGroup(c, p)
	c.JSON(resp, err)
}

func groupGetGroupInfo(c *bm.Context) {
	p := new(GroupInfoByIdReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GroupSvc.GetGroupInfo(c, p)
	c.JSON(resp, err)
}

func groupGetAllGroups(c *bm.Context) {
	p := new(SearchGroupReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GroupSvc.GetAllGroups(c, p)
	c.JSON(resp, err)
}

func groupGetAllGroupsByUid(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GroupSvc.GetAllGroupsByUid(c, p)
	c.JSON(resp, err)
}

func groupAddMember(c *bm.Context) {
	p := new(AddMemberReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GroupSvc.AddMember(c, p)
	c.JSON(resp, err)
}

// RegisterGroupBMServer Register the blademaster route
func RegisterGroupBMServer(e *bm.Engine, server GroupBMServer, midMap map[string]bm.HandlerFunc) {
	auth := midMap["auth"]
	GroupSvc = server
	e.POST("/group/create", auth, groupCreateGroup)
	e.GET("/group/info", groupGetGroupInfo)
	e.GET("/group/all", groupGetAllGroups)
	e.GET("/group/all/user", auth, groupGetAllGroupsByUid)
	e.POST("/group/addMember", auth, groupAddMember)
}
