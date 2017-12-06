package restapi

import (
	"github.com/zeebox/goose4"
	"net/http"
	"strings"
)

// Middleware handles the "/service" prefix to comply with the SE4 prefix
type Middleware struct {
	handler http.Handler
	SE4     goose4.Goose4
}

// NewMiddleware takes an http handler
// to wrap and returns mutable Middleware object
func NewMiddleware(h http.Handler) *Middleware {
	return &Middleware{
		handler: h,
	}
}

// ServeHTTP wraps our requests and handles any calles to `/service*`.
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.String(), "/service") {
		m.SE4.ServeHTTP(w, r)
	} else {
		m.handler.ServeHTTP(w, r)
	}
}
