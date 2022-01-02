package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
}

type Post struct {
	gorm.Model
	AuthorID uint
	Author   User
	Title    string `gorm:"not null;unique"`
	Content  string `gorm:"not null;"`
}
