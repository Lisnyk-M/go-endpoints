package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	// CreditCards []CreditCard
}

// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

type RegisterForm struct {
	User     string `json:"user" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
