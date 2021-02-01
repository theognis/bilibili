package tool

import "github.com/gin-gonic/gin"

//false则token为空，返回NO_TOKEN_PROVIDED
func CheckTokenNil(ctx *gin.Context, token string) bool {
	if token == "" {
		Failed(ctx, "NO_TOKEN_PROVIDED")
		return false
	}
	return true
}

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, gin.H{
		"status": true,
		"data":   v,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, gin.H{
		"status": false,
		"data":   v,
	})
}
