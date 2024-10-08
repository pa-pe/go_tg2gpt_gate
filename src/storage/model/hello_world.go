package model

import (
	"gorm.io/gorm"
	"upserv/src/storage/utils/auto_migrate"
)

type HelloWorld struct {
	gorm.Model
	Title string `gorm:"column:title"`
}

func init() {
	auto_migrate.RegisterModel(&HelloWorld{})
}
