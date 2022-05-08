package main

import (
	"fmt"
	"Task-Manager/utils"
	"Task-Manager/interface/models"
	
	"github.com/spf13/viper"
)

func main(){
	// 加载本地Config文件
	utils.LoadConfig()
	fmt.Println(viper.GetStringMap("postgres"))

	// 数据库连接测试
	issueDB := utils.GetIssueDBConnection()

	// 数据库Migrate
	// issueDB.AutoMigrate(&models.Issue{}, &models.IssueTag{}, &models.Tag{}, &models.User{}, &models.Comment{})

	// 插入数据
	issueDB.Create(&models.Issue{Title: "test", Content: "test"})
}