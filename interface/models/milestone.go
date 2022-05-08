package models

import (
	"gorm.io/gorm"
)

type Milestone struct {
	gorm.Model
	// Milestone主体内容
	Title		string
	Content		string
}