package service

import (
	"context"
	"upserv/src/service/cache"
	"upserv/src/service/internal"
	"upserv/src/storage"
	"upserv/src/storage/model"
)

type Services struct {
	HelloWorld  IHelloWorld
	TelegramBot ITelegramBot
}

type IHelloWorld interface {
	Find(ctx context.Context) (*model.HelloWorld, error)
}

type ITelegramBot interface {
	Start()
}

func NewServices(storages *storage.Storages, cache cache.ICache) *Services {
	return &Services{
		HelloWorld: internal.NewHelloWorldService(storages.HelloWorld, cache),
	}
}
