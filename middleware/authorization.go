package middleware

import (
	"log"
	"net/http"
)

var Auth = map[string]string{"vova": "vovapass", "seva": "sevapass"}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("basic auth status:", ok)
			return
		}

		if pass, ok := Auth[username]; !ok || pass != password {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("error authorization")
			return
		}
		next.ServeHTTP(w, r)
	})
}
