package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const RequstIDKey = "requestID"

// LogginMiddleware добавляет уникальный ID запроса к контексту и журналирует данные запроса.

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), RequstIDKey, requestID)
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))

		// Журналируем запрос
		log.Printf("Request ID: %s, Method: %s, URI: %s, Client IP: %s, Duration: %s\n",
			requestID, r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}
