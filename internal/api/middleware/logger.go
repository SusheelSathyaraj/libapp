package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//Before Handler
		log.Printf("%s %s", r.Method, r.RequestURI)

		//call handler
		next.ServeHTTP(w, r)

		//after handler
		log.Printf("%s completed in %v", r.RequestURI, time.Since(start))
	})
}
