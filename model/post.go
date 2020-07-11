package model

import (
	"github.com/jinzhu/gorm"
)

// Post is Entity posts
type Post struct {
	gorm.Model        // ID, created_at, updated_at, deleted_at
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Author     string `gorm:"not null"`
}
