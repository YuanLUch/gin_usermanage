package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm:"type:varchar(255);not null;unique" json:"username"`
	Password    string `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber string `gorm:"type:varchar(255);not null" json:"phone_number"`
	Email       string `gorm:"type:varchar(255);not null" json:"email"`
	Role        int    `gorm:"type:int;not null;DEFAULT:1" json:"role"`
}
