package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"
)

func CreateDevice(conn *sql.DB, name, room, location, deviceId, address, make, model string) error {
	cmd := fmt.Sprintf("INSERT INTO devices (name, room, location, device_id, address, make, model) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s');", name, room, location, deviceId, address, make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Create Device: ", err)
	}
	return err
}

func ReadDevice(conn *sql.DB, name string) (*models.Device, error) {
	cmd := fmt.Sprintf("SELECT * FROM devices WHERE name = '%s';", name)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Println("DB Read Device: ", err)
		return nil, err
	}
	defer rows.Close()

	var device models.Device
	for rows.Next() {
		err := rows.Scan(&device.Name, &device.Room, &device.Location, &device.DeviceId, &device.Address, &device.Make, &device.Model)
		if err != nil {
			log.Println("DB Read Device: ", err)
			return nil, err
		}
	}
	return &device, nil
}

func ReadAllDevices(conn *sql.DB) (*[]models.Device, error) {
	rows, err := conn.Query("SELECT * FROM devices;")
	if err != nil {
		log.Println("DB Read All Devices: ", err)
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		err := rows.Scan(&device.Name, &device.Room, &device.Location, &device.DeviceId, &device.Address, &device.Make, &device.Model)
		if err != nil {
			log.Print("DB Read All Devices: ", err)
			return nil, err
		}
		devices = append(devices, device)
	}
	return &devices, nil
}

func ReadDevicesByMake(conn *sql.DB, make string) (*[]models.Device, error) {
	cmd := fmt.Sprintf("SELECT * FROM devices WHERE make = '%s';", make)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Println("DB Read Devices: ", err)
		return nil, err
	}
	defer rows.Close()

	var devices []models.Device
	for rows.Next() {
		var device models.Device
		err := rows.Scan(&device.Name, &device.Room, &device.Location, &device.DeviceId, &device.Address, &device.Make, &device.Model)
		if err != nil {
			log.Print("DB Read Devices: ", err)
			return nil, err
		}
		devices = append(devices, device)
	}
	return &devices, nil
}

// refactor to include all fields, nil check?
func updateDeviceKeyValue(conn *sql.DB, make, model, key, value string) error {
	cmd := fmt.Sprintf("UPDATE devices SET '%s' = '%s' WHERE make = '%s' AND model = '%s';", key, value, make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Update Device: ", err)
	}
	return err
}

func DeleteDevice(conn *sql.DB, name string) error {
	cmd := fmt.Sprintf("DELETE FROM devices WHERE name = '%s';", name)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Delete Device: ", err)
	}
	return err
}
