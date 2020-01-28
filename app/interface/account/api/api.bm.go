// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package api

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathAccountGetBasicInfo = "/account/basicInfo"

// AccountBMServer is the server API for Account service.
type AccountBMServer interface {
	GetBasicInfo(ctx context.Context, req *BasicInfoRequest) (resp *BasicInfo, err error)
}

var AccountSvc AccountBMServer

func accountGetBasicInfo(c *bm.Context) {
	p := new(BasicInfoRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AccountSvc.GetBasicInfo(c, p)
	c.JSON(resp, err)
}

// RegisterAccountBMServer Register the blademaster route
func RegisterAccountBMServer(e *bm.Engine, server AccountBMServer, midMap map[string]bm.HandlerFunc) {
	AccountSvc = server
	e.GET("/account/basicInfo", accountGetBasicInfo)
}
