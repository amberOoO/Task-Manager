package tagView

import (
	"Task-Manager/storage/models"
	"Task-Manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateTag(c *gin.Context) {
	var (
		json	CreateTagMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}
	// 创建tag
	tag := models.Tag{Content: json.Content}
	tagService := service.NewTagService()
	// 创建tag异常处理
	if err := tagService.CreateTag(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": tag.ToJson()})
}