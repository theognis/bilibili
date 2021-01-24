package main

import (
    "bilibili/controller"
    "bilibili/tool"
    "github.com/gin-gonic/gin"
    "html/template"
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
    space := router.Group("space")
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
    space.GET("/:uid", func(c *gin.Context){
        uid := c.Param("uid")
        c.HTML(http.StatusOK, "space.html", gin.H{
            "uid": uid,
            "tab": template.HTML(`<div id="tab_home" class="now_tab">主页</div>
            <div id="tab_moments">动态</div>
            <div id="tab_post">投稿</div>`),
        })
    })
    space.GET("/:uid/moments", func(c *gin.Context){
        uid := c.Param("uid")
        c.HTML(http.StatusOK, "space.html", gin.H{
            "uid": uid,
            "tab": template.HTML(`<div id="tab_home">主页</div>
            <div id="tab_moments" class="now_tab">动态</div>
            <div id="tab_post">投稿</div>`),
        })
    })
    space.GET("/:uid/post", func(c *gin.Context){
        uid := c.Param("uid")
        c.HTML(http.StatusOK, "space.html", gin.H{
            "uid": uid,
            "tab": template.HTML(`<div id="tab_home">主页</div>
            <div id="tab_moments">动态</div>
            <div id="tab_post" class="now_tab">投稿</div>`),
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