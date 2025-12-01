package config

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"smart_house/be/server"
)

type Message struct {
	Text string `json:"text"`
}

func Serve(conn *sql.DB) {
	// hello config
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(Message{Text: "hello config API"})
		}
	})

	// login/logout
	http.HandleFunc("/config/login", func(w http.ResponseWriter, r *http.Request) {
		server.Login(w, r, conn)
	})
	http.HandleFunc("/config/logout", func(w http.ResponseWriter, r *http.Request) {
		server.Logout(w, r, conn)
	})

	// users
	http.HandleFunc("/config/users", func(w http.ResponseWriter, r *http.Request) {
		server.UserHandler(w, r, conn)
	})

	// device models
	http.HandleFunc("/config/device-models", func(w http.ResponseWriter, r *http.Request) {
		server.DeviceModelHandler(w, r, conn)
	})

	log.Println("Config Server Started: 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
