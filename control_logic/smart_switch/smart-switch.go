package smart_switch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"smart_house/be/control_logic"
	"smart_house/be/db/models"
	"strings"

	"github.com/gorilla/websocket"
)

type SmartSwitch struct {
	device         *models.Device
	output         bool
	status         *Status
	OnStatusChange func(string, *Status)
}

func NewSmartSwitch(device *models.Device, onStatusChange func(string, *Status)) control_logic.Controllable {
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
	return []string{"SetSwitch", "ToggleSwitch", "GetSwitchStatus", "ConnectWebsocket", "DisconnectWebsocket"}
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

// WEBSOCKET METHODS
func (s *SmartSwitch) parseWebsocketMessage(msg string) {
	var wsp WebSocketResponse
	jsonReader := strings.NewReader(msg)
	decoder := json.NewDecoder(jsonReader)
	err := decoder.Decode(&wsp)
	if err != nil {
		log.Printf("Smart Switch %s ParseWebsocketMessage: %s\n", msg, err)
	}

	fmt.Printf("Smart Switch %s ParseWebsocketMessage: %+v\n", s.device.Name, wsp)
	fmt.Printf("Smart Switch %s Output: %t\n", s.device.Name, wsp.Result.SwitchZero.Output)
}

func (s *SmartSwitch) ConnectWebsocket() error {
	//url := "ws://" + s.device.Address + "/rpc"
	u := url.URL{Scheme: "ws", Host: s.device.Address, Path: "/rpc"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("Smart Switch %s ConnectWebsocket: %s\n:", s.device.Room, err)
		return err
	}

	err = c.WriteJSON(WebSocketRequest{Id: s.device.DeviceId, Src: "smart-house", Method: "Shelly.GetStatus"})
	if err != nil {
		log.Println("Smart Switch %s ConnectWebsocket Write: %s\n:", s.device.Room, err)
		return err
	}

	// create channel to signal interrupt
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Smart Switch %s ConnectWebsocket Read: %s\n", s.device.DeviceId, err)
				return
			}
			s.parseWebsocketMessage(string(message))
		}
	}()
	return nil
}
