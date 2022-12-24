package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model `swaggerignore:"true"`
	Email    string `gorm:"unique"`
	Password string
}