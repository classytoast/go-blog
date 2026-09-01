package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(
	handler http.Handler,
	mids ...Middleware,
) http.Handler {
	if len(mids) == 0 {
		return handler
	}

	for i := len(mids) - 1; i >= 0; i-- {
		handler = mids[i](handler)
	}

	return handler
}
