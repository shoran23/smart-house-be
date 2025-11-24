package smart_switch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"smart_house/be/db/models"
)

type SmartSwitch struct {
	device         *models.Device
	output         bool
	status         *Status
	OnStatusChange func(string, *Status)
}

func NewSmartSwitch(device *models.Device, onStatusChange func(string, *Status)) *SmartSwitch {
	return &SmartSwitch{
		device:         device,
		output:         false,
		status:         nil,
		OnStatusChange: onStatusChange,
	}
}

func (s *SmartSwitch) GetDeviceInfo() *models.Device {
	return s.device
}

func (s *SmartSwitch) GetControlMethods() []string {
	return []string{"SetSwitch", "ToggleSwitch", "GetSwitchStatus"}
}

// SWITCH METHODS
func handleSwitch(url string) (*SwitchState, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var swState SwitchState
	err = json.Unmarshal(body, &swState)
	if err != nil {
		return nil, err
	}
	return &swState, nil
}

func (s *SmartSwitch) SetSwitch(state bool) {
	url := fmt.Sprintf("http://%s/rpc/Switch.Set?id=%d&on=%t", s.device.Address, s.device.DeviceId, state)
	swState, err := handleSwitch(url)
	if err != nil {
		log.Fatalf("Smart Switch %s SetSwitch: %s", s.device.DeviceId, err)
	}
	if swState.WasOn != state {
		s.output = state
	}
}

func (s *SmartSwitch) ToggleSwitch() {
	url := fmt.Sprintf("http://%s/rpc/Switch.Toggle?id=%d", s.device.Address, s.device.DeviceId)
	swState, err := handleSwitch(url)
	if err != nil {
		log.Fatalf("Smart Switch %s ToggleSwitch: %s", s.device.DeviceId, err)
	}
	s.output = !swState.WasOn
}

func (s *SmartSwitch) GetSwitchStatus() {
	url := fmt.Sprintf("http://%s/rpc/Switch.GetStatus?id=%d", s.device.Address, s.device.DeviceId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Smart Switch %s GetSwitchStatus: %s", s.device.DeviceId, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Smart Switch %s GetSwitchStatus: %s", s.device.DeviceId, err)
	}

	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Fatalf("Smart Switch %s GetSwitchStatus: %s", s.device.DeviceId, err)
	}
	s.status = &status
	s.output = status.Output
	if s.OnStatusChange != nil {
		s.OnStatusChange(s.device.DeviceId, s.status)
	}
}
