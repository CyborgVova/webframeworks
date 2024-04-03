package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := slog.With(
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("remote_address", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()))

		next.ServeHTTP(w, r)

		start := time.Now()
		log.Info("", slog.String("time_duration", time.Since(start).String()))
	})
}
