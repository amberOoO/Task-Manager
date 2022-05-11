package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	// Content主要内容
	Title		string
	Content		string
	// 外键
	CreaterID	uint
	IssueID		uint
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id": c.ID,
		"title": c.Title,
		"content": c.Content,
		"createrID": c.CreaterID,
		"issueID": c.IssueID,
	}
}