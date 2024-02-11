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
	// 主页
	i.IndexRouterGroup.GET("/index", i.IndexService.GetIndexPage)
	// 社区
	i.IndexRouterGroup.GET("/community", i.IndexService.GetCommunityPage)
	// 公告
	i.IndexRouterGroup.GET("/announcement", i.IndexService.GetAnnouncementPage)
	// 发布
	i.IndexRouterGroup.GET("/posting", i.IndexService.GetPostingPage)

	i.IndexRouterGroup.GET("/posting/postTopic", i.IndexService.GetPostingTopicPage)
	// 关注
	i.IndexRouterGroup.GET("/concern", i.IndexService.GetConcernPage)
}
