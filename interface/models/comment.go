package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	// Content主要内容
	Content		string
	// 关联的Issue
	IssueID		uint
	// 创建者ID
	CreaterID	uint
}