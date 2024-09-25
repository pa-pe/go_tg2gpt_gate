package src

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"upserv/config"
	_ "upserv/docs"
	"upserv/logger"
	"upserv/src/apperror"
	"upserv/src/http/middleware"
	"upserv/src/http/request"
	"upserv/src/http/response"
)

type Routes []Route

// ROUTE
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  []middleware.I
}

func (route Route) handle(w http.ResponseWriter, r *http.Request) {
	//apply middleware
	if route.Middleware != nil {
		var err *apperror.IError
		for _, v := range route.Middleware {
			w, r, err = v.Run(w, r)
			if err != nil {
				logger.Log.Debug(err.Error())
				response.Error(err, w)
				return
			}
		}
	}
	route.HandlerFunc.ServeHTTP(w, r)
}

func NewRouter(middlewares *middleware.MiddlewareImp) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	routes := newRouteList(middlewares)
	//Default health check route
	router.Path("/health").
		Methods("GET").
		Name("HealthCheck").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//TODO:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		})

	// Swagger docs
	router.PathPrefix("/swagger/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		httpSwagger.WrapHandler.ServeHTTP(w, r)
	})
	for _, route := range routes {
		methods := []string{route.Method}
		router.
			Methods(methods...).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.handle)
	}

	return logRequest(defaultHeaders(router))
}

func defaultHeaders(handler http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", config.RequestIdKey})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT", "DELETE"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := uuid.New()
			r.Header.Set(config.RequestIdKey, id.String())
			w.Header().Set(config.RequestIdKey, id.String())
			w.Header().Set("host-env", getEnv())
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			request.ExtractLang(r)
			ctx := context.WithValue(r.Context(), config.RequestIdKey, id.String())
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		}))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ignore options logging
		if r.Method == "OPTIONS" || r.URL.Path == "/health" {
			handler.ServeHTTP(w, r)
			return
		}
		l := logger.NewRequestLog(w, r)
		defer l.Commit()
		handler.ServeHTTP(l.GetResponseWriter(), l.GetRequest())
	})
}

// Simple helper function to read an environment or return a default value
func getEnv() string {
	s := "unrecognized"
	if value, err := os.Hostname(); err == nil {
		s = value
	}
	return s
}
