package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smart_house/be/db"
)

func getDeviceModels(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	deviceModels, err := db.ReadAllDeviceModels(conn)
	if err != nil {
		http.Error(w, "Error reading device models", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(deviceModels)
	if err != nil {
		http.Error(w, "Error encoding device models", http.StatusInternalServerError)
		return
	}
}

func getDeviceModel(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	devMake := r.URL.Query().Get("make")
	devModel := r.URL.Query().Get("model")
	if devMake == "" {
		http.Error(w, "Missing device make query", http.StatusBadRequest)
		return
	}
	if devModel == "" {
		http.Error(w, "Missing device model query", http.StatusBadRequest)
		return
	}
	deviceModel, err := db.ReadDeviceModel(conn, devMake, devModel)
	if err != nil {
		http.Error(w, "Error reading device model", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(deviceModel)
}

func handleGetDeviceModels(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	params := r.URL.Query()
	if len(params) > 0 {
		getDeviceModel(w, r, conn)
	} else {
		getDeviceModels(w, r, conn)
	}
}

func postDeviceModel(w http.ResponseWriter, r *http.Request, conn *sql.DB) {

}

func DeviceModelHandler(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		handleGetDeviceModels(w, r, conn)
	case http.MethodPost:
		postDeviceModel(w, r, conn)
	}
}
