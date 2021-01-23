package model

import "github.com/dgrijalva/jwt-go"

type MyCustomClaims struct {
	Name string `json:"name"`
	phone string `json:"email"`
	jwt.StandardClaims
}
