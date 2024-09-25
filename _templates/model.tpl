package model

import "gorm.io/gorm"

type {{ .ModelName }} struct {
	gorm.Model
	Title string `gorm:"column:title"`
}
