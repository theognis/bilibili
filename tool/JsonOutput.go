package tool

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, v interface{})  {
	ctx.JSON(200, gin.H{
		"status": true,
		"data": v,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, gin.H{
		"status": false,
		"data": v,
	})
}
