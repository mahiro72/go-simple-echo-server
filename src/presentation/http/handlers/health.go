package handlers

import (
	"net/http"

	"github.com/mahiro72/go-simple-echo-server/presentation/http/response"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Ping(w http.ResponseWriter, r *http.Request) {
	response.WriteJson(w, ping{Message: "hello"}, http.StatusOK)
}

type ping struct {
	Message string `json:"message"`
}
