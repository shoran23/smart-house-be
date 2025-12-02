package main

import (
	"fmt"
	"log"
	"smart_house/be/db"
	"smart_house/be/runtime"
	"smart_house/be/server/config"
	"smart_house/be/server/control"
)

func main() {
	fmt.Println("Hello Smart House")

	// connect database
	conn, err := db.Connect()
	if err != nil {
		log.Fatal("Unable to connect to Database: ", err)
	}

	// create default user
	err = db.CreateDefaultUser(conn)
	if err != nil {
		log.Fatal("Unable to create default user: ", err)
	}

	// initialize runtimes
	dr := runtime.InitializeDevices(conn)
	devices := dr.GetAllDevices()
	fmt.Println("Devices: ", len(devices))
	for _, d := range devices {
		fmt.Printf("Device: %+v ", d.GetDeviceInfo())
	}

	// launch servers
	go control.Serve(conn, dr)
	go config.Serve(conn)
	// launch the ui server
	select {}
}
