package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	Userinfo Userinfo
	Type     string //"REFRESH_TOKEN"表示为一个refresh token，"TOKEN"表示为一个token
	Time     time.Time
	jwt.StandardClaims
}
