package config

import (
	"gitlab.com/AngelX/common/config"
	"time"
)

const (
	FirewallRequestsPerMinute = 1000
	RequestIdKey              = "X-Request-Id"
	FirewallBanPeriod         = 10 * time.Second
)

func Get(section string, key string) string {
	return config.Get(section, key)
}

func GetInt(section string, key string) int {
	return config.GetInt(section, key)
}

func GetBool(section string, key string) bool {
	return config.GetBool(section, key)
}
