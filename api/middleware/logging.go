package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"Method: %s, Endpoint: %s, Response Time: %s",
			r.Method,
			r.URL.Path,
			time.Since(startTime),
		)
	})
}
