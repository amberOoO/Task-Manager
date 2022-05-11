package main

import (
	// "Task-Manager/storage/models"
	"Task-Manager/utils"
	"Task-Manager/view/userView"
	"Task-Manager/view/issueView"
	"Task-Manager/view/milestoneView"
	"Task-Manager/view/commentView"
	"Task-Manager/view/tagView"
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

	// 启动Gin
	router := gin.Default()

	/******* user *******/
	// 创建用户接口
	router.POST("/user/create", userView.CreateUser)

	/******* issue *******/
	// 创建issue接口
	router.POST("/issue/create", issueView.CreateIssue)
	// 更新issue接口
	router.POST("/issue/update", issueView.UpdateIssue)
	// 更新issueTags接口
	router.POST("/issue/updateTags", issueView.UpdateIssueTags)
	// 更新issueMilestone接口
	router.POST("/issue/updateMilestone", issueView.UpdateIssueMilestone)
	// 删除issue接口
	router.POST("/issue/delete", issueView.DeleteIssue)
	// 分页获取issue接口
	router.POST("/issue/getIssueWithCondition", issueView.GetIssueWithCondition)

	/******* milestone *******/
	// 创建milestone接口
	router.POST("/milestone/create", milestoneView.CreateMilestone)
	// 更新milestone接口
	router.POST("/milestone/update", milestoneView.UpdateMilestone)
	// 删除milestone接口
	router.POST("/milestone/delete", milestoneView.DeleteMilestone)

	/******* comment *******/
	// 创建comment接口
	router.POST("/comment/create", commentView.CreateComment)

	/******* tag *******/
	// 创建tag接口
	router.POST("/tag/create", tagView.CreateTag)

	router.Run("0.0.0.0:8888")
}