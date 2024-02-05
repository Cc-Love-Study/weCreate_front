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
