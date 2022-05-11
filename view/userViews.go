package view

import (
	"Task-Manager/storage/models"
	"Task-Manager/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateUser(c *gin.Context) {
	var(
		json	CreateUserMsg
	)
	// 绑定json数据
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "data format error"})
		return
	}
	// 创建用户
	user := models.User{Nickname: json.Nickname}
	userService := service.NewUserService()

	if err := userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	// 返回成功结果
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": user.ToJson()})
}