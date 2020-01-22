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

var PathPassportRegister = "/passport/register"
var PathPassportLogin = "/passport/login"

// PassportBMServer is the server API for Passport service.
type PassportBMServer interface {
	Register(ctx context.Context, req *RegisterReq) (resp *RegisterResp, err error)

	// `midware:"basic"`
	Login(ctx context.Context, req *google_protobuf1.Empty) (resp *LoginResp, err error)
}

var PassportSvc PassportBMServer

func passportRegister(c *bm.Context) {
	p := new(RegisterReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSvc.Register(c, p)
	c.JSON(resp, err)
}

func passportLogin(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := PassportSvc.Login(c, p)
	c.JSON(resp, err)
}

// RegisterPassportBMServer Register the blademaster route
func RegisterPassportBMServer(e *bm.Engine, server PassportBMServer, midMap map[string]bm.HandlerFunc) {
	basic := midMap["basic"]
	PassportSvc = server
	e.POST("/passport/register", passportRegister)
	e.GET("/passport/login", basic, passportLogin)
}
