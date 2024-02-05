package controllers

import (
	"weCreate/services"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	IndexRouterGroup *gin.RouterGroup
	IndexService     *services.IndexService
}

// 工厂函数
func NewIndexController(r *gin.Engine, name string, indexService *services.IndexService) *IndexController {
	rGroup := r.Group("/" + name)
	return &IndexController{IndexRouterGroup: rGroup, IndexService: indexService}
}

// 路由注册
func (i *IndexController) InitIndexController() {
	i.IndexRouterGroup.GET("/index", i.IndexService.GetIndexPage)
}
