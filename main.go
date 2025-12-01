package main

import (
	"fmt"
	"log"
	control_logic "smart_house/be/control-logic"
	"smart_house/be/control-logic/smart_switch"
	"smart_house/be/db"
	"smart_house/be/server/config"
	"smart_house/be/server/control"
	"smart_house/be/state"
)

func main() {
	fmt.Println("Hello Smart House")

	// initialize state
	state.InitDeviceState()

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

	// pull control device data and
	devices, err := db.ReadAllDevices(conn)
	if err != nil {
		log.Fatal("Unable to read devices: ", err)
	}
	for _, d := range *devices {
		fmt.Println("Device Name: ", d.Name)
		// need to get model of device
		dm, err := db.ReadDeviceModel(conn, d.Make, d.Model)
		if err != nil {
			log.Fatalf("Unable to read %s device model: %s", d.Name, err)
		}
		switch dm.DeviceType {
		case int(control_logic.SmartSwitch):
			ss := smart_switch.NewSmartSwitch(&d, nil)
		}
	}

	// update the state based on the device data

	// launch servers
	go control.Serve(conn)
	go config.Serve(conn)
	// launch the ui server
	select {}
}
