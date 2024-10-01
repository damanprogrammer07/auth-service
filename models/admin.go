package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"unique" json:"password"`
}
