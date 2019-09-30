package sample

import (
	"encoding/json"
	"net/http"
)

// OkHandler returns 200
func OkHandler(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Status string `json:"status"`
	}
	body := Body{Status: "OK"}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
