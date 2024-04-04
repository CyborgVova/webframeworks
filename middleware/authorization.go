package middleware

import (
	"log"
	"net/http"
	"webframeworks/storage"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("basic auth status:", ok)
			return
		}

		if pass, ok := storage.Auth[username]; !ok || pass != password {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("error authorization")
			return
		}
		next.ServeHTTP(w, r)
	})
}
