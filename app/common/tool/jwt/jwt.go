/*
@Time : 2020/1/21 22:36
@Author : Minus4
*/
package jwt

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

var jwtSecret string

func Init(secret string) {
	jwtSecret = secret
}

type Claims struct {
	Uid        int64
	DeviceId   string
	DeviceType int
	jwt.StandardClaims
}

const (
	WebClient = iota + 1
	AppleMobileClient
	AndroidClient
	PcClient
	AppleMacClient
)

type AuthParams struct {
	Uid        int64
	DeviceId   string
	DeviceType int
}

func GenerateToken(p AuthParams) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 36)

	claims := Claims{
		p.Uid,
		// TODO use real device id
		uuid.NewV4().String(),
		p.DeviceType,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "minus4.cn",
			Id:        uuid.NewV4().String(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func ParseBearerHeader(header string) (token string) {
	s := strings.SplitN(header, " ", 2)
	if len(s) != 2 {
		return
	} else {
		authType := s[0]
		if authType != "Bearer" {
			return ""
		}
		token = s[1]
		return
	}
}
