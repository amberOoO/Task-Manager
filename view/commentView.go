package view

import (
	"Task-Manager/storage/models"
	"Task-Manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"fmt"
)

func CreateComment(c *gin.Context) {
	var (
		json	CreateCommentMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}
	fmt.Println(json)

	comment := models.Comment{
		Title: json.Title, 
		Content: json.Content,
		IssueID: json.IssueID,
		CreaterID: json.CreaterID, 
	}
	commentService := service.NewCommentService()
	// 创建comment异常处理
	if err := commentService.CreateComment(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	fmt.Println(comment)
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": comment.ToJson()})
}