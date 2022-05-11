package view

import (
	"Task-Manager/service"
	"Task-Manager/storage/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func CreateIssue(c *gin.Context) {
	var (
		json	CreateIssueMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}

	issue := models.Issue{
		Title: json.Title, 
		Content: json.Content,
		CreaterID: json.CreaterID, 
		MilestoneID: json.MilestoneID,
		Tags: models.StringArrToTagArr(json.Tags),
	}
	issueService := service.NewIssueService()
	// 创建issue异常处理
	if err := issueService.CreateIssue(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": issue.ToJson()})
}

func UpdateIssue(c *gin.Context) {
	var (
		json	UpdateIssueMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}

	issue := models.Issue{
		Model: gorm.Model{
			ID: json.ID,
		},
		Content: json.Content,
		CreaterID: json.CreaterID, 
		MilestoneID: json.MilestoneID,
	}

	issueService := service.NewIssueService()
	// 创建issue异常处理
	if err := issueService.UpdateIssue(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": issue.ToJson()})
}

func UpdateIssueTags(c *gin.Context) {
	var (
		json	UpdateIssueTagsMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}

	issue := models.Issue{
		Model: gorm.Model{
			ID: json.ID,
		},
		Tags: models.StringArrToTagArr(json.Tags),
	}

	issueService := service.NewIssueService()
	// 创建issue异常处理
	if err := issueService.UpdateIssueTags(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": issue.ToJson()["tags"]})
}

func UpdateIssueMilestone(c *gin.Context) {
	var (
		json	UpdateIssueMilestoneMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}

	issue := models.Issue{
		Model: gorm.Model{
			ID: json.ID,
		},
	}
	issueService := service.NewIssueService()
	// 调用issue服务，创建issue异常处理
	if err := issueService.UpdateMilestone(&issue, json.MilestoneID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": issue.ToJson()})
}

func DeleteIssue(c *gin.Context) {
	var (
		json	DeleteIssueMsg
	)
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}

	issue := models.Issue{
		Model: gorm.Model{
			ID: json.ID,
		},
	}

	issueService := service.NewIssueService()
	// 创建issue异常处理
	if err := issueService.DeleteIssue(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": issue.ToJson()["tags"]})
}