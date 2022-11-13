package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);"`
	Email    string `gorm:"type:varchar(100);unique;"`
	Password string `gorm:"type:varchar(100);"`
}
