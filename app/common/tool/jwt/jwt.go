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

var JwtSecret string

func Setup(jwtSecret string) {
	JwtSecret = jwtSecret
}

type Claims struct {
	// user id
	Uid          int64
	RefreshTimes int
	jwt.StandardClaims
}

func GenerateToken(id int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 4)

	claims := Claims{
		id,
		0,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "minus4.cn",
			Id:        uuid.NewV4().String(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(JwtSecret))

	return token, err
}

func RefreshToken(claims Claims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
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
