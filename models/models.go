package models

import "gorm.io/gorm"

// basic
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
}

// a blog post
type Post struct {
	gorm.Model
	AuthorId uint `json:"author_id"`
	Author   User `gorm:"embedded;embeddedPrefix:author_;foreignKey:AuthorId json:"author"`
}
