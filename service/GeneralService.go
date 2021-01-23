package service

import (
	"bilibili/model"
	"bilibili/tool"
	"github.com/dgrijalva/jwt-go"
)

//通用的服务
type GeneralService struct {

}

//解析Token
func (u *UserService) ParseToken(tokenString string) (*model.MyCustomClaims, error) {
	JwtCfg := tool.GetCfg().Jwt
	mySigningKey := []byte(JwtCfg.SigningKey)
	token, err := jwt.ParseWithClaims(tokenString, &model.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if clams, ok := token.Claims.(*model.MyCustomClaims); ok && token.Valid {
		return clams, nil
	} else {
		return nil, err
	}
}

//构建一个jwt

//更新jwt

//...

