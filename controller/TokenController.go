package controller

import (
	"bilibili/service"
	"bilibili/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TokenController struct {
}

func (g *TokenController) Router(engine *gin.Engine) {
	engine.GET("/api/verify/token", g.getToken)
}

//通过refreshToken获取token
func (g *TokenController) getToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refreshToken")

	gs := service.TokenService{}
	us := service.UserService{}

	//判断refreshToken状态
	model, err := gs.ParseRefreshToken(refreshToken)
	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(ctx, "refreshToken失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return

	}

	if model.Type == "ERR" {
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return
	}

	//根据username更新userinfo
	newUserinfo, err := us.GetUserinfo(model.Userinfo.Username)

	//创建新token
	newToken, err := gs.CreateToken(newUserinfo, 120, "TOKEN")
	if err != nil {
		fmt.Println("getTokenCreateErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	tool.Success(ctx, newToken)
}
