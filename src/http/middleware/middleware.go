package middleware

import (
	"net/http"
	"upserv/src/apperror"
	"upserv/src/http/middleware/internal"
	"upserv/src/service"
)

type MiddlewareImp struct {
	FirewallMiddleware *internal.Firewall
}

type I interface {
	Run(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, *apperror.IError)
}

func NewMiddlewares(services *service.Services) *MiddlewareImp {
	return &MiddlewareImp{
		FirewallMiddleware: internal.NewFirewallMiddleware(),
	}
}
