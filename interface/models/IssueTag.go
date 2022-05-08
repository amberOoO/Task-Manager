package models

import (
	"gorm.io/gorm"
)

type IssueTag struct {
	gorm.Model
	// IssueTag存的俩外键
	IssueID		uint
	TagID		uint
}