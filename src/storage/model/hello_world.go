package model

import "gorm.io/gorm"

type HelloWorld struct {
	gorm.Model
	Title string `gorm:"column:title"`
}
