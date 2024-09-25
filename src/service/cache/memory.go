package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"reflect"
	"sync"
	"time"
	"upserv/config"
)

func NewInMemoryCache() ICache {
	return &cacheData{
		engine: cache.New(15*time.Minute, 10*time.Minute),
	}
}

type cacheData struct {
	engine  *cache.Cache
	mu      sync.RWMutex
	KeyList map[string][]string
}

type listContainer struct {
	Data  interface{}
	Total int64
}

// Put a value in cache or replace existing
func (c *cacheData) Put(ctx context.Context, namespace string, key string, value interface{}, period time.Duration) {
	c.PutList(ctx, namespace, key, value, 1, period)
}

// PutList a value as a list with total in cache or replace existing
func (c *cacheData) PutList(ctx context.Context, namespace string, key string, value interface{}, total int64, period time.Duration) {
	if c.isActive() {
		err := c.putKeyToNamespace(ctx, namespace, key)
		if err == nil {
			c.engine.Set(c.buildKeyName(ctx, namespace, key), listContainer{Data: value, Total: total}, period)
		}
	}
}

// Load Extract cache to dest value or an error if key doesn't exist
func (c *cacheData) Load(ctx context.Context, namespace string, key string, dest interface{}) error {
	var total *int64
	return c.LoadList(ctx, namespace, key, dest, total)
}

// LoadList Extract cache to dest value with total or an error if key doesn't exist
func (c *cacheData) LoadList(ctx context.Context, namespace string, key string, dest interface{}, total *int64) error {
	if c.isActive() {
		value, found := c.engine.Get(c.buildKeyName(ctx, namespace, key))
		if !found {
			return errors.New(fmt.Sprintf("key %s is not found", key))
		}
		if value == nil {
			return errors.New(fmt.Sprintf("key %s is not found", key))
		}
		err := c.execute(dest, total, value)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("cache was disabled")
}

func (c *cacheData) Delete(ctx context.Context, namespace string, key string) {
	c.engine.Delete(c.buildKeyName(ctx, namespace, key))
}

func (c *cacheData) ClearNamespace(ctx context.Context, namespace string) {
	if c.isActive() {
		for _, key := range c.getNamespaceKeys(namespace) {
			c.engine.Delete(c.buildKeyName(ctx, namespace, key))
		}
		c.dropNamespaceKeys(ctx, namespace)
	}
}

func (c *cacheData) dropNamespaceKeys(ctx context.Context, namespace string) {
	c.mu.Lock()
	delete(c.KeyList, namespace)
	c.mu.Unlock()
}

func (c *cacheData) getNamespaceKeys(namespace string) []string {
	c.mu.RLock()
	val, ok := c.KeyList[namespace]
	c.mu.RUnlock()
	if !ok {
		return []string{}
	}
	return val
}

func (c *cacheData) buildKeyName(ctx context.Context, namespace string, key string) string {
	return "__namespace__" + namespace + "__key__" + key
}

func (c *cacheData) execute(dest interface{}, total *int64, value interface{}) error {
	reflectDest := reflect.ValueOf(dest)
	if !reflectDest.IsValid() {
		return errors.New("invalid destination")
	}
	if reflectDest.Kind() == reflect.Ptr {
		reflectDest = reflectDest.Elem()
	}

	reflectVal := reflect.ValueOf(value)
	if !reflectVal.IsValid() {
		return errors.New("invalid value")
	}

	totalStruct := reflect.Indirect(reflectVal).FieldByName("Total")
	if totalStruct.IsValid() && total != nil {
		*total = totalStruct.Int()
	}

	dataStruct := reflectVal.FieldByName("Data")
	if dataStruct.IsValid() {
		dataElem := dataStruct.Elem()
		if dataElem.Type() == reflectDest.Type() {
			reflectDest.Set(dataStruct.Elem())
		}
	}

	return nil
}

func (c *cacheData) isActive() bool {
	return config.GetBool("cache", "activated")
}

func (c *cacheData) putKeyToNamespace(ctx context.Context, namespace string, key string) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}

	c.mu.Lock()
	if c.KeyList == nil {
		c.KeyList = make(map[string][]string)
	}
	c.KeyList[namespace] = append(c.KeyList[namespace], key)
	c.mu.Unlock()

	return nil
}
