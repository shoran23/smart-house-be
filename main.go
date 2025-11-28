package main

import (
	"fmt"
	"log"
	"smart_house/be/db"
	"smart_house/be/db/models"
	"smart_house/be/server/config"
	"smart_house/be/server/control"
	"smart_house/be/state"
)

func main() {
	fmt.Println("Hello Smart House")

	// initialize state
	state.InitDeviceState()

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Unable to connect to Database: ", err)
		return
	}
	db.CreateUser(conn, "admin", "admin", models.AdminAccess)

	session, _ := db.ReadSession(conn, "scotty")
	if session == nil {
		log.Print("DB Read Session Not Found")
	} else {
		log.Print("DB Read Session Found")

	}

	// update the state

	// launch servers
	go control.Serve(conn)
	go config.Serve(conn)
	// launch the ui server
	select {}
}
