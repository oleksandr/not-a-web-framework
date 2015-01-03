package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Welcome!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "OK")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	json.NewEncoder(w).Encode(user)
}
