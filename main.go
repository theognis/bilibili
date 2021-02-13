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
	router.Static("statics", "./view/statics")
	account := router.Group("account")
	video := router.Group("video")
	space := router.Group("space")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./view/favicon.ico")
	})
	account.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "act_register.html", nil)
	})
	account.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "act_login.html", nil)
	})
	account.GET("/setting", func(c *gin.Context) {
		c.HTML(http.StatusOK, "act_setting.html", nil)
	})
	account.GET("/set/password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "act_setPassword.html", nil)
	})
	account.GET("/set/email", func(c *gin.Context) {
		c.HTML(http.StatusOK, "act_setEmail.html", nil)
	})
	account.GET("/set/phone", func(c *gin.Context) {
		c.HTML(http.StatusOK, "act_setPhone.html", nil)
	})
	video.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "video.html", gin.H{})
	})
	video.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "video_upload.html", gin.H{})
	})
	space.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "space.html", gin.H{})
	})

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
