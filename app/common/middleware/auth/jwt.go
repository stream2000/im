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
	"time"
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

type BearerConfig struct {
	JwtSecret      string
	MaxRefreshTime string
}

func BearerAuth(jwtSecret string, maxRefreshTime int) func(*bm.Context) {
	jwtTool.Setup(jwtSecret)
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
		// this is a unexpired token
		ctx.Set("uid", claims.Uid)

		now := time.Now()

		expiredTime := time.Unix(claims.ExpiresAt, 0)
		if expiredTime.Sub(now) < time.Minute*20 && claims.RefreshTimes < maxRefreshTime {
			claims.RefreshTimes += 1
			claims.ExpiresAt = now.Add(time.Hour * 2).Unix()
			newToken, _ := jwtTool.RefreshToken(*claims)
			ctx.Writer.Header().Set("Refresh-Token", newToken)
		}
	}
}
