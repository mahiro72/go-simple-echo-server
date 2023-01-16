package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any, status int) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(status)
}
