package state

import (
	"smart_house/be/control-logic"
	"sync"
)

type DeviceState struct {
	mu            sync.RWMutex
	controllables []*control_logic.Controllable
}

var DevState DeviceState

func InitDeviceState() {
	DevState = DeviceState{
		mu:            sync.RWMutex{},
		controllables: make([]*control_logic.Controllable, 0),
	}
}

func (s *DeviceState) AddAppliance(appliance *control_logic.Controllable) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.controllables = append(s.controllables, appliance)
}

func (s *DeviceState) GetAllDevices() []*control_logic.Controllable {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.controllables
}

func (s *DeviceState) GetAllDevicesByRoom(room string) []*control_logic.Controllable {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var found []*control_logic.Controllable
	//for _, c := range s.controllables {
	//	if c.) == room {
	//		found = append(found, a)
	//	}
	//}
	return found
}

func (s *DeviceState) GetDeviceByRoomLocation(room string, location string) *control_logic.Controllable {
	s.mu.RLock()
	defer s.mu.RUnlock()

	//for _, a := range s.controllables {
	//	if a.GetRoom() == room && a.GetLocation() == location {
	//		return a
	//	}
	//}
	return nil
}
