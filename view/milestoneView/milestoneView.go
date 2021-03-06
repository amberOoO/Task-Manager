package milestoneView

import (
	"Task-Manager/service"
	"Task-Manager/storage/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
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
	milestone := models.Milestone{Title: json.Title, Content: json.Content, Issues: json.GetIssueArr()}
	milestoneService := service.NewMilestoneService()

	if err := milestoneService.CreateMilestone(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": milestone.ToJson()})
}

func UpdateMilestone(c *gin.Context) {
	var (
		json				UpdateMilestoneMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}
	// 更新milestone
	milestone := models.Milestone{
		Model: gorm.Model{
			ID: json.ID,
		},
		Title: json.Title,
		Content: json.Content,
	}
	milestoneService := service.NewMilestoneService()

	if err := milestoneService.UpdateMilestone(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": milestone.ToJson()})
}

func DeleteMilestone(c *gin.Context) {
	var (
		json				UpdateMilestoneMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}

	// 删除milestone
	milestone := models.Milestone{
		Model: gorm.Model{
			ID: json.ID,
		},
	}
	milestoneService := service.NewMilestoneService()

	if err := milestoneService.DeleteMilestone(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": milestone.ToJson()})
}
