package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName	string
	// 其他一些属性，待定。如邮箱等
}