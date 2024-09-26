package logmiddleware

import (
	"log"
	"net/http"
	"time"
)

type StateRecorder struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader captures the status code and writes it to the underlying ResponseWriter
func (rec *StateRecorder) WriteHeader(statusCode int) {
    rec.StatusCode = statusCode
    rec.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleware logs the method, path, status code, and duration of each request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Wrap the ResponseWriter to capture the status code
		wrappedWriter := &StateRecorder{
			ResponseWriter: w,
			StatusCode: http.StatusOK,
		}

		// Pass the request to the next handler
		next.ServeHTTP(wrappedWriter, r)

		// Log the request details
		log.Printf("Method: %s, Path: %s, Status: %d, Duration: %s", r.Method, r.URL.Path, wrappedWriter.StatusCode, time.Since(start))
	})
}