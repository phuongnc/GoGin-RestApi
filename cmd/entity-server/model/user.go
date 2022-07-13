package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique, Index:idx_email;"`
	Password  string
	FirstName string
	LastName  string
}
