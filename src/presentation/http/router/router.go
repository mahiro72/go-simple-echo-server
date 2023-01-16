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
