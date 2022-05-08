package storageuploadscripts

import (
	"Task-Manager/interface/models"
	"Task-Manager/utils"
)

func InsertData() {
	issueDB := utils.GetIssueDBConnection()
	// 创建用户
	issueDB.Create(&models.User{NickName: "testUser1"})
	issueDB.Create(&models.User{NickName: "testUser2"})
	// 插入Issue数据
	issueDB.Create(&models.Issue{Title: "issue1", Content: "issue1 content"})
	issueDB.Create(&models.Issue{Title: "issue2", Content: "issue2 content"})
	issueDB.Create(&models.Issue{Title: "issue3", Content: "issue3 content"})
	// 插入Comment数据
	issueDB.Create(&models.Comment{Content: "comment1", IssueID: 1, CreaterID: 1})
}