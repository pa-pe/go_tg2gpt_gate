package internal

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"upserv/config"
	"upserv/logger"
	"upserv/src/apperror"
	"upserv/src/util"
)

type Firewall struct {
	LimitCount int
	Period     time.Duration
	BanPeriod  time.Duration
}

func (a *Firewall) Run(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, *apperror.IError) {
	ip := getIP(r)
	ctx := r.Context()
	key := fmt.Sprintf("%p%s", a, ip)
	if a.BanPeriod <= 0 {
		a.BanPeriod = config.FirewallBanPeriod
	}
	count := util.Cache.FindInt(key, 0)
	if count >= a.LimitCount {
		t := time.Until(util.Cache.TimeLeftForKey(key)).Truncate(time.Second)
		if t > a.BanPeriod {
			util.Cache.Put(key, count, a.BanPeriod)
			t = a.BanPeriod.Truncate(time.Second)
		}
		logger.Log.Debug("requests limit reached: ", count, "from ip", ip, "key is", key)
		logger.Log.Debug("Time left in cache: ", t)
		return w, r, apperror.Forbidden(ctx).WithMsg(fmt.Sprintf("Forbidden by firewall. to many connections, retry in %s", t))
	}
	count++
	util.Cache.Put(key, strconv.Itoa(count), a.Period)
	return w, r, nil
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func NewFirewallMiddleware() *Firewall {
	return &Firewall{
		LimitCount: config.FirewallRequestsPerMinute,
		Period:     1 * time.Minute,
	}
}
