package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("cc")
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("../static", "static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "begin  weCreate project front"})
	})
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.Run(":8080")
}
