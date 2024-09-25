package storage

import (
	"context"
	"gorm.io/gorm"
	"upserv/src/storage/internal"
	"upserv/src/storage/model"
)

type Storages struct {
	HelloWorld IHelloWorld
}

type IHelloWorld interface {
	Find(ctx context.Context) (*model.HelloWorld, error)
}

func NewStorages(db *gorm.DB) *Storages {
	return &Storages{
		HelloWorld: internal.NewHelloWorldStorage(db),
	}
}
