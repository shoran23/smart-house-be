package configserver

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func Serve() {
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(Message{Text: "hello config API"})
		}
	})
	log.Println("Config Server Started: 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
