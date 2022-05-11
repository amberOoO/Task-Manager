package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	// Tag主体内容
	Content		string		`gorm:"primarykey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	// 多对多关系
	Issues		[]Issue		`gorm:"many2many:issue_tags;"`
}

// 将Tag字符串列表转换为Tag列表
func StringArrToTagArr(tags []string) []Tag {
	var tagArr []Tag = make([]Tag, len(tags))
	for index, tag := range tags {
		tagArr[index] = (Tag{Content: tag})
	}
	return tagArr
}

// 将Tag列表转换为Tag字符串列表
func TagArrToStringArr(tags []Tag) []string {
	var tagArr []string = make([]string, len(tags))
	for index, tag := range tags {
		tagArr[index] = tag.Content
	}
	return tagArr
}

func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"content": t.Content,
	}
}