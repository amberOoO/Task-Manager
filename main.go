package main

import (
	// "Task-Manager/storage/models"
	"Task-Manager/utils"
	"Task-Manager/view"
	"fmt"

	// "Task-Manager/storage/models"
	// "Task-Manager/storageInsert"
	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func main(){
	// 加载本地Config文件
	utils.LoadConfig()
	fmt.Println(viper.GetStringMap("postgres"))

	// // 数据库连接测试
	// issueDB := utils.GetTaskDBConnection()
	// var issue1 = models.Issue{}
	// var comments []models.Comment
	// issue1.ID = 7
	// issueDB.Find(&issue1)
	// issueDB.Debug().Joins("left join issues on comments.issue_id = issues.id").Where("issues.id = ?", 7).Find(&comments)
	// fmt.Println(comments[0])

	// 数据库Migrate
	// issueDB.AutoMigrate(&models.Issue{})
	// issueDB.AutoMigrate(&models.Tag{})
	// issueDB.AutoMigrate(&models.User{})
	// issueDB.AutoMigrate(&models.Comment{})
	// issueDB.AutoMigrate(&models.Milestone{})


	// 插入数据
	// issueDB.Create(&models.Issue{Title: "test", Content: "test"})
	// storageInsert.StorageInsert()

	// 启动Gin
	router := gin.Default()

	// 创建用户接口
	router.POST("/createUser", view.CreateUser)

	// 创建issue接口
	router.POST("/createIssue", view.CreateIssue)
	// 更新issue接口
	router.POST("/updateIssue", view.UpdateIssue)
	// 更新issueTags接口
	router.POST("/updateIssueTags", view.UpdateIssueTags)
	// 删除issue接口
	router.POST("/deleteIssue", view.DeleteIssue)

	// 创建comment接口
	router.POST("/createComment", view.CreateComment)

	// 创建tag接口
	router.POST("/createTag", view.CreateTag)

	// 创建milestone接口
	router.POST("/createMilestone", view.CreateMilestone)
	// 更新milestone接口
	router.POST("/updateMilestone", view.UpdateMilestone)
	// 删除milestone接口
	router.POST("/deleteMilestone", view.DeleteMilestone)

	router.Run("0.0.0.0:8888")
}