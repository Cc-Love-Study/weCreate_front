package services

import (
	"net/http"
	"weCreate/utils"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	Utils *utils.Utils
}

// 工厂函数
func NewUserService(utils *utils.Utils) *UserService {
	return &UserService{Utils: utils}
}

func (u *UserService) GetLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
	return
}

func (u *UserService) PostLoginInfo(c *gin.Context) {
	var (
		email    string
		password string
	)
	email = c.PostForm("email")
	password = c.PostForm("password")
	// 路由转发 到后端 登录函数 接收后端返回内容 完成登录 如果登录成功给前端设置cookie
	c.JSON(http.StatusOK, u.Utils.ReturnSucess(200, "登录消息完成", gin.H{"email": email, "password": password}, 0))
	return
}
