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

// need to rework the appliances model structure
// in a database store the data for each device type
// ex: Shelly 1PM Gen4
// create id's for which class is used
// that class will require specific data to work
// the class data can be stored in db
// appliance data can also be stored in db

func main() {
	fmt.Println("Hello Smart House")

	// initialize state
	state.InitDeviceState()

	// create appliance devices
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
	// launch the control server
	// launch the config server
	// launch the ui server
	select {}
}
