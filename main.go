package main

import (
	"bilibili/controller"
	"bilibili/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := tool.GetCfg()

	router := gin.Default()
	routerRegister(router)

	if cfg.AppHttps {
		if err := router.RunTLS(cfg.AppHost+":"+cfg.AppPort, "config/ssl/anonym.ink_chain.crt", "config/ssl/anonym.ink_key.key"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := router.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
			log.Fatal(err)
		}
	}
}

func routerRegister(engine *gin.Engine) {
	new(controller.UserController).Router(engine)
	new(controller.TokenController).Router(engine)
}
