package services

import (
	"net/http"
	"weCreate/utils"

	"github.com/gin-gonic/gin"
)

type IndexService struct {
	Utils *utils.Utils
}

// 工厂函数
func NewIndexService(utils *utils.Utils) *IndexService {
	return &IndexService{Utils: utils}
}
func (u *IndexService) GetIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
	return
}

func (u *IndexService) GetCommunityPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "community.html", gin.H{})
	return
}

func (u *IndexService) GetAnnouncementPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "announcement.html", gin.H{})
	return
}

func (u *IndexService) GetPostingPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "posting.html", gin.H{})
	return
}

func (u *IndexService) GetPostingTopicPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "postwrite.html", gin.H{})
	return
}

func (u *IndexService) GetConcernPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "concern.html", gin.H{})
	return
}
