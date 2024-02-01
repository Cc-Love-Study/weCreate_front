package main

import (
	"fmt"
	"net/http"
	"os"
	"weCreate/conf"
	"weCreate/controllers"
	"weCreate/services"
	"weCreate/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

var appConf = conf.NewAppConf()

func init() {
	err := ini.MapTo(appConf, "./conf/conf.ini")
	if err != nil {
		fmt.Printf("ini文件读取失败:%v", err)
		os.Exit(1)
	}
	fmt.Println("当前配置:", appConf)
}
func main() {
	// 初始化路由
	r := gin.Default()
	// 加载资源
	r.LoadHTMLGlob("views/*")
	r.Static("../static", "static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "begin  weCreate project front"})
	})

	util := utils.NewUtils(appConf.Md5Key)
	// 用户部分 注册 登录路由页面路由
	userService := services.NewUserService(util)
	userController := controllers.NewUserController(r, "", userService)
	userController.InitUserController()

	r.Run(appConf.RunConf.Address)
}
