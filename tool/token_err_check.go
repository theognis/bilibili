package tool

import (
	"bilibili/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

//检查并解析token
//token过期返回TOKEN_EXPIRED
//token不正确返回PARSE_TOKEN_ERROR
func CheckTokenErr(ctx *gin.Context, claims *model.MyCustomClaims, err error) bool {
	if err == nil && claims.Type == "ERR" {
		Failed(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	if err != nil {
		fmt.Println("HAHA")
		if err.Error()[:16] == "token is expired" {
			Failed(ctx, "TOKEN_EXPIRED")
			return false
		}

		fmt.Println("getTokenParseTokenErr:", err)
		Failed(ctx, "PARSE_TOKEN_ERROR")
		return false
	}

	return true
}
