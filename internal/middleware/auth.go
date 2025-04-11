package middleware

import "net/http"

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Validate JWT here
		next.ServeHTTP(w, r)
	})
}
