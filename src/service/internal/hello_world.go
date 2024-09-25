package internal

import (
	"context"
	"fmt"
	"time"
	"upserv/src/service/cache"
	"upserv/src/storage"
	"upserv/src/storage/model"
)

type helloWorldImpl struct {
	cache             cache.ICache
	cacheTime         time.Duration
	cachePrefix       string
	helloWorldStorage storage.IHelloWorld
}

func (c *helloWorldImpl) cacheKey(methodName string, params interface{}) string {
	return fmt.Sprintf("__"+c.cachePrefix+"_%s_params_%+v", methodName, params)
}

func (c *helloWorldImpl) Find(ctx context.Context) (*model.HelloWorld, error) {
	helloWorld := &model.HelloWorld{}

	namespace := "helloWorld"
	key := "first"

	err := c.cache.Load(ctx, namespace, key, helloWorld)
	if err != nil {
		helloWorld, err = c.helloWorldStorage.Find(ctx)
		if err == nil && helloWorld != nil {
			c.cache.Put(ctx, namespace, key, *helloWorld, c.cacheTime)
		}
	}
	return helloWorld, err

}

func NewHelloWorldService(helloWorldStorage storage.IHelloWorld, cache cache.ICache) *helloWorldImpl {
	return &helloWorldImpl{
		cache:             cache,
		helloWorldStorage: helloWorldStorage,
		cacheTime:         5 * time.Minute,
		cachePrefix:       "HelloWorld",
	}
}
