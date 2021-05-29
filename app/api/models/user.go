package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Uuid     string `gorm:"unique" json:"uuid"`
}
