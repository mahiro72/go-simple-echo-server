package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	mux  *chi.Mux
	port int
}

type IRouter interface {
	setMiddleware()
	Serve() error
}

func New(port int) IRouter {
	r := &Router{
		mux:  chi.NewRouter(),
		port: port,
	}
	return r
}

func (r *Router) setMiddleware() {
	r.mux.Use(middleware.Logger)
	r.mux.Use(middleware.Recoverer)
	r.mux.Use(writeContentTypeJson)
}

func writeContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func (r *Router) Serve() error {
	// middlewareの設定
	r.setMiddleware()

	// health routerの初期化
	r.initHealth()

	return http.ListenAndServe(
		fmt.Sprintf(":%d", r.port),
		r.mux,
	)
}
