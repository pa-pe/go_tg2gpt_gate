package src

import (
	"upserv/src/handler"
	"upserv/src/http/middleware"
)

func newRouteList(middlewares *middleware.MiddlewareImp) []Route {
	return []Route{
		{
			Name:        "GetHelloWorld",
			Method:      "GET",
			Pattern:     "/hello_world",
			HandlerFunc: handler.GetHelloWorld,
			Middleware:  []middleware.I{middlewares.FirewallMiddleware},
		},
	}
}
