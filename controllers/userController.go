package controllers

import (
	"weCreate/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserRouterGroup *gin.RouterGroup
	UserService     *services.UserService
}

// 工厂函数
func NewUserController(r *gin.Engine, name string, userService *services.UserService) *UserController {
	rGroup := r.Group("/" + name)
	return &UserController{UserRouterGroup: rGroup, UserService: userService}
}

// 路由注册
func (u *UserController) InitUserController() {
	u.UserRouterGroup.GET("/register", u.UserService.GetRegisterPage)
	u.UserRouterGroup.POST("/register/post", u.UserService.PostRegisterInfo)
	u.UserRouterGroup.GET("/login", u.UserService.GetLoginPage)
	u.UserRouterGroup.POST("/login/post", u.UserService.PostLoginInfo)
}
