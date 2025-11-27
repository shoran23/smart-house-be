package main

import (
	"fmt"
	"log"
	"smart_house/be/configserver"
	"smart_house/be/controlserver"
	"smart_house/be/db"
	"smart_house/be/db/models"
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

	// update the state

	// launch servers
	go controlserver.Serve()
	go configserver.Serve()
	// launch the ui server
	select {}
}
