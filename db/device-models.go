package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"
)

func CreateDeviceModel(conn *sql.DB, make, model, purpose string, deviceType int) error {
	cmd := fmt.Sprintf("INSERT INTO device_models (make, model, purpose, device_type) VALUES ('%s', '%s', '%s', '%d');", make, model, purpose, deviceType)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Print("DB Create Device Model: ", err)
	}
	return err
}

func ReadDeviceModel(conn *sql.DB, make, model string) (*models.DeviceModel, error) {
	cmd := fmt.Sprintf("SELECT * FROM device_models WHERE make='%s' AND model='%s';", make, model)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Println("DB Read Device Model: ", err)
		return nil, err
	}
	defer rows.Close()

	var deviceModel models.DeviceModel
	for rows.Next() {
		err := rows.Scan(&deviceModel.Make, &deviceModel.Model, &deviceModel.Purpose, &deviceModel.DeviceType)
		if err != nil {
			log.Println("DB Read Device Model: ", err)
			return nil, err
		}
	}
	return &deviceModel, nil
}

func ReadAllDeviceModels(conn *sql.DB) (*[]models.DeviceModel, error) {
	rows, err := conn.Query("SELECT * FROM device_models;")
	if err != nil {
		log.Println("DB Read Device Models: ", err)
		return nil, err
	}
	defer rows.Close()

	var deviceModels []models.DeviceModel
	for rows.Next() {
		var deviceModel models.DeviceModel
		err := rows.Scan(&deviceModel.Make, &deviceModel.Model, &deviceModel.Purpose, &deviceModel.DeviceType)
		if err != nil {
			log.Println("DB Read Device Models: ", err)
			return nil, err
		}
		deviceModels = append(deviceModels, deviceModel)
	}
	return &deviceModels, nil
}

// redo this update to include all values to update, nil check?
func updateDeviceModelKeyValue(conn *sql.DB, make, model, key, value string) error {
	cmd := fmt.Sprintf("UPDATE device_models SET '%s' = '%s' WHERE make='%s' AND model='%s';", key, value, make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Update Device Model: ", err)
	}
	return err
}

func DeleteDeviceModel(conn *sql.DB, make, model string) error {
	cmd := fmt.Sprintf("DELETE FROM device_models WHERE make='%s' AND model='%s';", make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Delete Device Model: ", err)
	}
	return err
}
