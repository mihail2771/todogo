package core_http_middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return h
	}
	for i := len(m) - 1; i >= 0; i-- {
		if m[i] == nil {
			continue
		}
		h = m[i](h)
	}
	return h
}
