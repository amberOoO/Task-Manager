package models

import (
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	// Issue主体内容
	Title			string
	Content			string
	// 外键
	CreaterID		uint
	MilestoneID		uint
	// 一对多关系
	Comments		[]Comment
	// 多对多关系
	Tags			[]Tag		`gorm:"many2many:issue_tags;"`
}

func (Issue) TableName() string {
	return "issues"
}

func (i *Issue) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id": i.ID,
		"title": i.Title,
		"content": i.Content,
		"createrID": i.CreaterID,
		"milestoneID": i.MilestoneID,
		"tags": TagArrToStringArr(i.Tags),
	}
}