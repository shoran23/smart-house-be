package control

import (
	"database/sql"
	"log"
	"net/http"
	"smart_house/be/server"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Serve(conn *sql.DB) {
	http.HandleFunc("/control-live", controlLiveHandler)
	http.HandleFunc("/control", controlHandler)
	http.HandleFunc("/control/appliances", controlApplianceHandler)
	http.HandleFunc("/control/login", func(w http.ResponseWriter, r *http.Request) {
		server.Login(w, r, conn)
	})
	log.Println("Control Server Started :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
