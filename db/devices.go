package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"
)

func CreateDevice(conn *sql.DB, name, room, location, deviceId, address, make, model string) {
	cmd := fmt.Sprintf("INSERT INTO devices (name, room, location, deviceId, address, make, model) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s');", name, room, location, deviceId, address, make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Create Device: ", err)
	}
}

func ReadDevice(conn *sql.DB, make, model string) *models.Device {
	cmd := fmt.Sprintf("SELECT * FROM device WHERE make = '%s' AND model = '%s';", make, model)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Fatal("DB Read Device: ", err)
		return nil
	}
	defer rows.Close()

	var device models.Device
	for rows.Next() {
		err := rows.Scan(&device.Make, &device.Model, &device.Location, &device.DeviceId, &device.Address, &device.Make, &device.Model)
		if err != nil {
			log.Fatal("DB Read Device: ", err)
			return nil
		}
	}
	return &device
}

func ReadAllDevices(conn *sql.DB) *[]models.Device {
	rows, err := conn.Query("SELECT * FROM devices;")
	if err != nil {
		log.Fatal("DB Read All Devices: ", err)
		return nil
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		err := rows.Scan(&device.Make, &device.Model, &device.Location, &device.DeviceId, &device.Address, &device.Make, &device.Model)
		if err != nil {
			log.Fatal("DB Read All Devices: ", err)
			return nil
		}
		devices = append(devices, device)
	}
	return &devices
}

func ReadDevicesByMake(conn *sql.DB, make string) *[]models.Device {
	cmd := fmt.Sprintf("SELECT * FROM devices WHERE make = '%s';", make)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Fatal("DB Read Devices: ", err)
		return nil
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		err := rows.Scan(device.Make, &device.Model, &device.Location, &device.DeviceId, &device.Address, &device.Make, &device.Model)
		if err != nil {
			log.Fatal("DB Read Devices: ", err)
			return nil
		}
		devices = append(devices, device)
	}
	return &devices
}

func UpdateDeviceKeyValue(conn *sql.DB, make, model, key, value string) {
	cmd := fmt.Sprintf("UPDATE devices SET '%s' = '%s' WHERE make = '%s' AND model = '%s';", key, value, make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Update Device: ", err)
		return
	}
}

func DeleteDevice(conn *sql.DB, make, model string) {
	cmd := fmt.Sprintf("DELETE FROM device WHERE make = '%s' AND model = '%s';", make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Delete Device: ", err)
		return
	}
}
