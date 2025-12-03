package runtime

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/control_logic"
	"smart_house/be/control_logic/smart_switch"
	"smart_house/be/db"
	"sync"
)

type DeviceRuntime struct {
	mu            sync.RWMutex
	controllables []control_logic.Controllable
}

var dr DeviceRuntime

func (s *DeviceRuntime) AddControllable(c control_logic.Controllable) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.controllables = append(s.controllables, c)
}

func (s *DeviceRuntime) GetAllDevices() []control_logic.Controllable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.controllables
}

func (s *DeviceRuntime) GetDevice() *control_logic.Controllable {
	return nil
}

func (s *DeviceRuntime) GetAllDevicesByRoom(room string) []control_logic.Controllable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var found []control_logic.Controllable
	for _, c := range s.controllables {
		d := c.GetDeviceInfo()
		if d.Room == room {
			found = append(found, c)
		}
	}
	return found
}

func (s *DeviceRuntime) GetDeviceByRoomLocation(room, location string) []control_logic.Controllable {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var found []control_logic.Controllable
	for _, c := range s.controllables {
		d := c.GetDeviceInfo()
		if d.Room == room && d.Location == location {
			found = append(found, c)
		}
	}
	return found
}

func InitializeDevices(conn *sql.DB) *DeviceRuntime {
	dr = DeviceRuntime{
		mu:            sync.RWMutex{},
		controllables: make([]control_logic.Controllable, 0),
	}

	// get devices from database
	devices, err := db.ReadAllDevices(conn)
	if err != nil {
		log.Fatal("Runtime Devices ReadAllDevices:", err)
	}
	for _, d := range *devices {

		dm, err := db.ReadDeviceModel(conn, d.Make, d.Model)

		if err != nil {
			log.Fatal("Runtime Devices ReadDeviceModel:", err)
		}

		fmt.Printf("dm = %+v\n", dm.DeviceType)

		switch dm.DeviceType {
		case int(control_logic.TypeSmartSwitch):
			fmt.Println("SmartSwitch Device Found")
			c := smart_switch.NewSmartSwitch(&d, nil)
			c.(*smart_switch.SmartSwitch).ConnectWebsocket()
			dr.AddControllable(c)
		}
	}

	return &dr
}
