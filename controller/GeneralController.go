package controller

import (
	"bilibili/service"
	"bilibili/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GeneralController struct {

}

func (g *GeneralController) Router(engine *gin.Engine) {
	engine.GET("/api/verify/token", g.getToken)
}

//通过refreshToken获取token
func (g *GeneralController) getToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refeshToken")

	gs := service.GeneralService{}

	//判断refreshToken状态
	model, err := gs.ParseToken(refreshToken)
	if err != nil {
		if err.Error()[:16] == "token is expired" {
			tool.Failed(ctx, "refreshToken失效")
			return
		}

		fmt.Println("getTokenParseTokenErr:", err)
		tool.Failed(ctx, "refreshToken不正确或系统错误")
		return

	}

	//创建新token
	newToken, err := gs.CreateToken(model.Userinfo, 120)
	if err != nil {
		fmt.Println("getTokenCreateErr:", err)
		tool.Failed(ctx, "系统错误")
		return
	}

	tool.Success(ctx, newToken)
}
