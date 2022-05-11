package view

import (
	"Task-Manager/storage/models"
	"Task-Manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateMilestone(c *gin.Context) {
	var(
		json	CreateMilestoneMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}
	// 创建milestone
	milestone := models.Milestone{Title: json.Title, Content: json.Content}
	milestoneService := service.NewMilestoneService()

	if err := milestoneService.CreateMilestone(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": milestone.ToJson()})
}