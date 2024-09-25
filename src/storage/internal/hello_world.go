package internal

import (
	"context"
	"gorm.io/gorm"
	"upserv/src/storage/model"
)

type helloWorldImpl struct {
	DB *gorm.DB
}

func (c *helloWorldImpl) Find(ctx context.Context) (*model.HelloWorld, error) {
	db := c.DB.WithContext(ctx)
	m := &model.HelloWorld{}
	db.Raw("Select 'Hello world' as title").Scan(m)

	return m, nil
}

func NewHelloWorldStorage(db *gorm.DB) *helloWorldImpl {
	return &helloWorldImpl{
		DB: db,
	}
}
