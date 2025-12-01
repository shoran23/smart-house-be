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

func (s *DeviceState) AddControllable(c *control_logic.Controllable) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.controllables = append(s.controllables, c)
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
	for _, controllable := range s.controllables {
		c := *controllable
		d := c.GetDeviceInfo()
		if d.Room == room {
			found = append(found, controllable)
		}
	}
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
