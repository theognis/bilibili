package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//检查并解析token
//token过期返回TOKEN_EXPIRED
//token不正确返回PRASE_TOKEN_ERROR
func CheckTokenErr(ctx *gin.Context, err error) bool {
	if err != nil {
		if err.Error()[:16] == "token is expired" {
			Failed(ctx, "TOKEN_EXPIRED")
			return false
		}

		fmt.Println("getTokenParseTokenErr:", err)
		Failed(ctx, "refreshToken不正确或系统错误")
		return false
	}

	return true
}
