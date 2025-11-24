package control_logic

import "smart_house/be/db/models"

type Controllable interface {
	GetDeviceInfo() *models.Device
	GetControlMethods() []string
}
