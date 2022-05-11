package models

import (
	"gorm.io/gorm"
)

type Milestone struct {
	gorm.Model
	// Milestone主体内容
	Title		string
	Content		string
	// 属于多个Issue
	Issues		[]Issue
}

func (Milestone) TableName() string {
	return "milestones"
}

func (m *Milestone) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id": m.ID,
		"title": m.Title,
		"content": m.Content,
	}
}