package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./view/html/*")
    router.Static("statics","./view/statics")
    passport := router.Group("passport")
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
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}