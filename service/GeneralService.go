package service

import (
	"bilibili/model"
	"bilibili/tool"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//通用的服务
type GeneralService struct {

}

//解析Token
func (u *GeneralService) ParseToken(tokenString string) (*model.MyCustomClaims, error) {
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
func (u *GeneralService) CreateToken(userinfo model.Userinfo, ExpireTime int64) (string, error) {
	JwtCfg := tool.GetCfg().Jwt
	mySigningKey := []byte(JwtCfg.SigningKey)

	claims := model.MyCustomClaims{
		Userinfo: userinfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + ExpireTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
//更新jwt

//...

