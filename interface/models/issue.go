package models

import (
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	// Issue主体内容
	Title		string
	Content		string
	// 附加的内容
	MilestoneID	uint
	// 创建者ID
	CreaterID	uint
}