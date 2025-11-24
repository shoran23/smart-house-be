package controlserver

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Text string `json:"text"`
}

func controlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message{Text: "Welcome to the Control Server"})
	}
}
