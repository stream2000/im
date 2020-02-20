/*
@Time : 2020/1/19 16:00
@Author : Minus4
*/
package auth

import (
	jwtTool "chat/app/common/tool/jwt"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/dgrijalva/jwt-go"
)

var (
	JwtTokenExpired = ecode.New(20000)
)

func init() {
	cm := map[int]string{
		JwtTokenExpired.Code(): "the json web token has expired",
	}
	ecode.Register(cm)
}

func BearerAuth(jwtSecret string, maxRefreshTime int) func(*bm.Context) {
	jwtTool.Init(jwtSecret)
	return func(ctx *bm.Context) {
		bearerHeader := ctx.Request.Header.Get("Authorization")
		token := jwtTool.ParseBearerHeader(bearerHeader)

		if token == "" {
			ctx.JSON(nil, ecode.Errorf(ecode.Unauthorized, " no bearer auth header in request headers"))
			ctx.Abort()
			return
		}
		claims, err := jwtTool.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ctx.JSON(nil, JwtTokenExpired)
				ctx.Abort()
				return
			default:
				ctx.JSON(nil, ecode.Errorf(ecode.Unauthorized, "error parse token"))
				ctx.Abort()
				return
			}
		}
		ctx.Set("uid", claims.Uid)
	}
}
