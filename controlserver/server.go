package controlserver

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Serve() {
	http.HandleFunc("/control-live", controlLiveHandler)
	http.HandleFunc("/control", controlHandler)
	http.HandleFunc("/control/appliances", controlApplianceHandler)
	log.Println("Control Server Started :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
