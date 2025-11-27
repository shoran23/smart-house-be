package main

import (
	"log"
	"smart_house/be/db"
)

//func main() {
//	fmt.Println("Hello Smart House")
//
//	// initialize state
//	state.InitDeviceState()
//
//	// create appliance devices
//	conn, err := db.Connect()
//	if err != nil {
//		log.Fatal("Unable to connect to Database: ", err)
//		return
//	}
//	db.CreateUser(conn, "admin", "admin", models.AdminAccess)
//
//	// update the state
//
//	// launch servers
//	go controlserver.Serve()
//	go configserver.Serve()
//	// launch the control server
//	// launch the config server
//	// launch the ui server
//	select {}
//}

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Unable to connect to Database: ", err)
		return
	}

	_ = db.DeleteSession(conn, "admin")

	select {}
}
