package internal

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"upserv/src/storage/model"
)

type helloWorldImpl struct {
	DB *gorm.DB
}

func (c *helloWorldImpl) Find(ctx context.Context) (*model.HelloWorld, error) {
	if c.DB == nil {
		err := errors.New("database is not initialized")
		return nil, err
	}

	db := c.DB.WithContext(ctx)
	m := &model.HelloWorld{}

	//	db.Raw("Select 'Hello world' as title").Scan(m)
	err := db.Raw("Select 'Hello world' as title").Scan(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func NewHelloWorldStorage(db *gorm.DB) *helloWorldImpl {
	return &helloWorldImpl{
		DB: db,
	}
}
