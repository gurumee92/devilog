package model

import "github.com/jinzhu/gorm"

// Account is Entity account
type Account struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"not null"`
	Picture  string `gorm:"not null"`
}
