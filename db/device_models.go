package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"
)

func CreateDeviceModel(conn *sql.DB, make, model, purpose string, deviceType int) {
	cmd := fmt.Sprintf("INSERT INTO device_models (make, model, purpose, device_type) VALUES ('%s', '%s', '%s', '%d');", make, model, purpose, deviceType)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Create Device Model: ", err)
	}
}

func ReadDeviceModel(conn *sql.DB, make, model string) *models.DeviceModel {
	cmd := fmt.Sprintf("SELECT * FROM device_models WHERE make='%s' AND model='%s';", make, model)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Fatal("DB Read Device Model: ", err)
		return nil
	}
	defer rows.Close()

	var deviceModel models.DeviceModel
	for rows.Next() {
		err := rows.Scan(&deviceModel.Make, &deviceModel.Model, &deviceModel.Purpose, &deviceModel.DeviceType)
		if err != nil {
			log.Fatal("DB Read Device Model: ", err)
			return nil
		}
	}
	return &deviceModel
}

func ReadAllDeviceModels(conn *sql.DB) *[]models.DeviceModel {
	rows, err := conn.Query("SELECT * FROM device_models;")
	if err != nil {
		log.Fatal("DB Read Device Models: ", err)
		return nil
	}
	defer rows.Close()

	var deviceModels []models.DeviceModel
	for rows.Next() {
		var deviceModel models.DeviceModel
		err := rows.Scan(&deviceModel.Make, &deviceModel.Model, &deviceModel.Purpose, &deviceModel.DeviceType)
		if err != nil {
			log.Fatal("DB Read Device Models: ", err)
			return nil
		}
		deviceModels = append(deviceModels, deviceModel)
	}
	return &deviceModels
}

func UpdateDeviceModelKeyValue(conn *sql.DB, make, model, key, value string) {
	cmd := fmt.Sprintf("UPDATE device_models SET '%s' = '%s' WHERE make='%s' AND model='%s';", key, value, make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Update Device Model: ", err)
	}
}

func DeleteDeviceModel(conn *sql.DB, make, model string) {
	cmd := fmt.Sprintf("DELETE FROM device_models WHERE make='%s' AND model='%s';", make, model)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Delete Device Model: ", err)
	}
}
