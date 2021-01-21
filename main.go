package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./view/html/*")
    router.Static("statics","./view/statics")
    router.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "index.html", nil)
    })
    router.GET("/favicon.ico", func(c *gin.Context){
        c.File("./view/favicon.ico")
    })
    router.Run(":8081")
}
