package utils

import (
	"golang.org/x/time/rate"
	"net/http"
)

// Limit set to 1000 requests per second with a burst of 5000.
var limiter = rate.NewLimiter(rate.Limit(1000), 5000)

// RateLimitMiddleware checks if the request is allowed by the rate limiter.
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
