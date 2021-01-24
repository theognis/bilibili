package model

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	Userinfo    Userinfo `json:"phone"`
	jwt.StandardClaims
}
