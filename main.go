package main

import (
    "bilibili/controller"
    "bilibili/tool"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func main() {
    cfg := tool.GetCfg()

    router := gin.Default()

    router.LoadHTMLGlob("./view/html/*")
    router.Static("statics","./view/statics")
    passport := router.Group("passport")
    video := router.Group("video")
    router.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "index.html", nil)
    })
    router.GET("/favicon.ico", func(c *gin.Context){
        c.File("./view/favicon.ico")
    })
    passport.GET("/register", func(c *gin.Context){
        c.HTML(http.StatusOK, "register_pc.html", nil)
    })
    passport.GET("/login", func(c *gin.Context){
        c.HTML(http.StatusOK, "login_pc.html", nil)
    })
    video.GET("/:av", func(c *gin.Context){
        av := c.Param("av")
        c.HTML(http.StatusOK, "video.html", gin.H{
            "av": av,
        })
    })

    routerRegister(router)

    if err := router.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
        log.Fatal(err)
    }
}

func routerRegister(engine *gin.Engine)  {
    new(controller.UserController).Router(engine)
}