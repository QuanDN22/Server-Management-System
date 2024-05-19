package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"column:username; required; unique; not null;"`
	Password string `json:"password" gorm:"column:password; required; not null;"`
	Email    string `json:"email" gorm:"column:email; required; unique; not null;"`
	Role     string `json:"role" gorm:"column:role; required; not null;"`
}
