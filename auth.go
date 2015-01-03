package main

import (
	"net/http"

	"github.com/gorilla/context"
)

func authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		context.Set(r, "user", map[string]string{"username": "Alex"})
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
