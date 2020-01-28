/*
@Time : 2020/1/22 16:09
@Author : Minus4
*/
package auth

import (
	"encoding/base64"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"strings"
)

func BasicFilter(ctx *bm.Context) {
	basicHeader := ctx.Request.Header.Get("Authorization")
	email, password, err := ParseBasicHeader(basicHeader)
	if err != nil {
		ctx.JSON(nil, err)
		ctx.Abort()
		return
	}
	ctx.Set("email", email)
	ctx.Set("password", password)
}

func ParseBasicHeader(header string) (email, password string, err error) {
	if header == "" {
		err = ecode.Errorf(ecode.Unauthorized, " no basic auth header in request headers")
		return
	}
	s := strings.SplitN(header, " ", 2)
	if len(s) != 2 {
		err = ecode.Errorf(ecode.Unauthorized, "the format of basic auth header is wrong")
		return
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		err = ecode.Errorf(ecode.Unauthorized, "the format of basic auth header is wrong")
		return
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 || pair[0] == "" || pair[1] == "" {
		err = ecode.Errorf(ecode.Unauthorized, "the content of basic auth header is wrong")
		return
	}
	email = pair[0]
	password = pair[1]
	return
}
