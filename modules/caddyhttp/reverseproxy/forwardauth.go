package reverseproxy

import (
	"net/http"
)

// ... existing code ...

func (f *ForwardAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.Handler) error {
	// ... existing auth request logic ...

	// After successful authentication from the auth service:
	// 1. Strip client-supplied headers that are managed by forward_auth
	for _, headerName := range f.CopyHeaders {
		r.Header.Del(headerName)
	}

	// 2. Copy headers from the auth response to the upstream request
	for _, headerName := range f.CopyHeaders {
		if val := authResponse.Header.Get(headerName); val != "" {
			r.Header.Set(headerName, val)
		}
	}

	return next.ServeHTTP(w, r)
}