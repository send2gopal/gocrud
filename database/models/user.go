package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	PasswordText string
	FirstName    string
	LastName     string
	Age          int16
}
