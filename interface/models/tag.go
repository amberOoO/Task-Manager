package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	// Tag主体内容
	Content		string
}