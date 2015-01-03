package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

func main() {
	var addr = flag.String("addr", ":8080", "HTTP bind address")

	flag.Parse()

	commonHandlers := alice.New(
		context.ClearHandler,
		loggingHandler,
		recoverHandler,
	)

	router := newRouter()
	router.get("/admin", commonHandlers.Append(authHandler).ThenFunc(adminHandler))
	router.get("/health", commonHandlers.ThenFunc(healthHandler))
	router.get("/", commonHandlers.ThenFunc(indexHandler))

	http.ListenAndServe(*addr, router)
}
