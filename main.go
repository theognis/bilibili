package main

import (
    "github.com/gin-gonic/gin"
    "pilipili/controller"
    "pilipili/tool"
)

func main() {
    cfg := tool.GetCfg()

    engine :=gin.Default()
    registerRouter(engine)

    engine.Run(cfg.AppPort + ":" + cfg.AppHost)
}

func registerRouter(engine *gin.Engine) {
    new(controller.UserController).Router(engine)
}
