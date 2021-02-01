package tool

import (
	"bilibili/model"
	"bilibili/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

//检查并解析token
//token为空返回NO_TOKEN_PROVIDED
//token过期返回TOKEN_EXPIRED
//token不正确返回PRASE_TOKEN_ERROR
func CheckTokenErr(ctx *gin.Context, token string) (model.Userinfo, bool) {
	if token == "" {
		Failed(ctx, "NO_TOKEN_PROVIDED")
		return model.Userinfo{}, false
	}

	gs := service.GeneralService{}
	claims, err := gs.ParseToken(token)

	if err != nil {
		if err.Error()[:16] == "token is expired" {
			Failed(ctx, "TOKEN_EXPIRED")
			return model.Userinfo{}, false
		}

		fmt.Println("getTokenParseTokenErr:", err)
		Failed(ctx, "refreshToken不正确或系统错误")
		return model.Userinfo{}, false
	}

	userinfo := claims.Userinfo
	return userinfo, true
}
