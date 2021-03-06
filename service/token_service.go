package service

import (
	"bilibili/model"
	"bilibili/tool"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//通用的服务
type TokenService struct {
}

//解析refreshToken
func (u *TokenService) ParseRefreshToken(tokenString string) (*model.MyCustomClaims, error) {
	JwtCfg := tool.GetCfg().Jwt
	mySigningKey := []byte(JwtCfg.SigningKey)
	token, err := jwt.ParseWithClaims(tokenString, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if clams, ok := token.Claims.(*model.MyCustomClaims); ok && token.Valid {
		if clams.Type == "TOKEN" {
			errClaims := new(model.MyCustomClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}
}

//解析Token
func (u *TokenService) ParseToken(tokenString string) (*model.MyCustomClaims, error) {
	JwtCfg := tool.GetCfg().Jwt
	mySigningKey := []byte(JwtCfg.SigningKey)
	token, err := jwt.ParseWithClaims(tokenString, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if clams, ok := token.Claims.(*model.MyCustomClaims); ok && token.Valid {
		if clams.Type == "REFRESH_TOKEN" {
			errClaims := new(model.MyCustomClaims)
			errClaims.Type = "ERR"
			return errClaims, nil
		}
		return clams, nil
	} else {
		return nil, err
	}
}

//构建一个jwt
func (u *TokenService) CreateToken(userinfo model.Userinfo, ExpireTime int64, tokenType string) (string, error) {
	JwtCfg := tool.GetCfg().Jwt
	mySigningKey := []byte(JwtCfg.SigningKey)

	claims := model.MyCustomClaims{
		Userinfo: userinfo,
		Type:     tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

//更新jwt

//...
