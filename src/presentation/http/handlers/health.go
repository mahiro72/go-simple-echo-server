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
	res := map[string]interface{}{
		"message": "hello",
	}
	response.WriteJson(w, res, http.StatusOK)
}
