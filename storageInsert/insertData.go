package storageInsert

import (
	"Task-Manager/storage/models"
	"Task-Manager/utils"
)

func StorageInsert() {
	issueDB := utils.GetTaskDBConnection()
	// 创建用户
	user1 := models.User{Nickname: "testUser1"}
	user2 := models.User{Nickname: "testUser2"}
	issueDB.Create(&user1)
	issueDB.Create(&user2)
	// 插入Issue数据
	issue1 := models.Issue{Title: "issue1", Content: "issue1 content", CreaterID: user1.ID}
	issue2 := models.Issue{Title: "issue2", Content: "issue2 content", CreaterID: user1.ID}
	issue3 := models.Issue{Title: "issue3", Content: "issue3 content", CreaterID: user2.ID}
	issueDB.Create(&issue1)
	issueDB.Create(&issue2)
	issueDB.Create(&issue3)
	// 插入Comment数据
	comment1 := models.Comment{Content: "comment1 content", IssueID: issue1.ID, CreaterID: user1.ID}
	issueDB.Create(&comment1)
}