package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/classytoast/go-blog/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type ApiVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []core_http_middleware.Middleware
}

func NewApiVersionRouter(
	apiVersion ApiVersion,
	middleware ...core_http_middleware.Middleware,
) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (vr *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		vr.Handle(pattern, route.WithMiddleware())
	}
}

func (vr *ApiVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(
		vr,
		vr.middleware...,
	)
}
