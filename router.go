package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type router struct {
	*httprouter.Router
}

// NewRouter is a router constructor
func newRouter() *router {
	return &router{httprouter.New()}
}

func (r *router) get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func (r *router) post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

func (r *router) put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

func (r *router) delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

func (r *router) patch(path string, handler http.Handler) {
	r.PATCH(path, wrapHandler(handler))
}

func (r *router) head(path string, handler http.Handler) {
	r.HEAD(path, wrapHandler(handler))
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}
