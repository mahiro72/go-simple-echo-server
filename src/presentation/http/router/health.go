package router

import "github.com/mahiro72/go-simple-echo-server/presentation/http/handlers"

func (r *Router) initHealth() {
	h := handlers.NewHealth()

	r.mux.Get("/ping", h.Ping)
}
