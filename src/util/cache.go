package util

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

var Cache cacheClient

var mc = cache.New(15*time.Minute, 10*time.Minute)

type cacheClient struct{}

// Put value in cache or replace existing
func (c *cacheClient) Put(key string, value interface{}, period time.Duration) {
	mc.Set(key, value, period)
}

// Return string value or error if key doesn't exist
func (c *cacheClient) Get(key string) (string, error) {
	str, found := mc.Get(key)
	if !found {
		return "", errors.New(fmt.Sprintf("key %s is not found", key))
	}
	return fmt.Sprintf("%v", str), nil
}

// Return string value if key found or default value
func (c *cacheClient) Find(key string, defaultValue string) string {
	v, err := c.Get(key)
	if err != nil {
		return defaultValue
	}
	return v
}

// Return int value or error if key doesn't exist
func (c *cacheClient) GetInt(key string) (int, error) {
	v, found := mc.Get(key)
	if !found {
		return 0, errors.New(fmt.Sprintf("key %s is not found", key))
	}
	val, err := strconv.Atoi(fmt.Sprintf("%v", v))
	if err != nil {
		return 0, nil
	}
	return val, nil
}

// Return int value if key found or default value
func (c *cacheClient) FindInt(key string, defaultValue int) int {
	v, err := c.GetInt(key)
	if err != nil {
		return defaultValue
	}
	return v
}

func (c *cacheClient) TimeLeftForKey(key string) time.Time {
	_, t, _ := mc.GetWithExpiration(key)
	return t
}
