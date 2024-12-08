package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Middleware для чтения и передачи сквозного идентификатора
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requsetID := r.Header.Get("X-Request-ID")
		if requsetID == "" {
			requsetID = generateNewRequestID()
		}

		// Добавляем идентификатор в контекст запроса
		ctx := context.WithValue(r.Context(), "requestID", requsetID)
		r = r.WithContext(ctx)

		// Добавляем идентификатор в заголовок ответа
		w.Header().Set("X-Request-ID", requsetID)

		// Логируем начало запроса
		log.Printf("Started %s %s, RequestID: %s", r.Method, r.URL.Path, requsetID)

		//Продолжение выполнения запроса
		next.ServeHTTP(w, r)
	})
}

func generateNewRequestID() string {
	return "REQ-" + strconv.FormatInt(time.Now().UnixNano(), 36)
}

// Middleware для логирования запросов.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логгируем время начала и дургие детали запроса
		log.Printf("Request started: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		// Логируем время завершения запроса.
		log.Printf("Request complited: %s %s, Duration: %v", r.Method, r.URL.Path, time.Since(start))

	})
}
