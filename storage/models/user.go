package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname	string
	// User对应多个Issue和Comment
	Comments	[]Comment	`gorm:"foreignKey:CreaterID"`
	Issues		[]Issue		`gorm:"foreignKey:CreaterID"`
	// 其他一些属性，待定。如邮箱等
}

func (User) TableName() string {
	return "users"
}

func (u *User) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id": u.ID,
		"Nickname": u.Nickname,
	}
}