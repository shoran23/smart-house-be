package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smart_house/be/db"
	"smart_house/be/db/models"
)

func getDeviceModels(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	dm, err := db.ReadAllDeviceModels(conn)
	if err != nil {
		http.Error(w, "Error reading device models", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dm)
	if err != nil {
		http.Error(w, "Error encoding device models", http.StatusInternalServerError)
		return
	}
}

func getDeviceModel(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	dMake := r.URL.Query().Get("make")
	dModel := r.URL.Query().Get("model")
	if dMake == "" {
		http.Error(w, "Missing device make query", http.StatusBadRequest)
		return
	}
	if dModel == "" {
		http.Error(w, "Missing device model query", http.StatusBadRequest)
		return
	}
	deviceModel, err := db.ReadDeviceModel(conn, dMake, dModel)
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
	var dm models.DeviceModel
	if r.Body == nil {
		http.Error(w, "Missing device model body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&dm)
	if err != nil {
		http.Error(w, "Error decoding device model body", http.StatusBadRequest)
		return
	}
	err = db.CreateDeviceModel(conn, dm.Make, dm.Model, dm.Purpose, dm.DeviceType)
	if err != nil {
		http.Error(w, "Error creating device model", http.StatusInternalServerError)
		return
	}
}

func putDeviceModel(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var dm models.DeviceModel
	if r.Body == nil {
		http.Error(w, "Missing device model body", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dm)
	if err != nil {
		http.Error(w, "Error decoding device model body", http.StatusBadRequest)
		return
	}
	err = db.CreateDeviceModel(conn, dm.Make, dm.Model, dm.Purpose, dm.DeviceType)
	if err != nil {
		http.Error(w, "Error creating device model", http.StatusInternalServerError)
		return
	}
}

func patchDeviceModel(w http.ResponseWriter, r *http.Request, conn *sql.DB) {

}

func deleteDeviceModel(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	dMake := r.URL.Query().Get("make")
	dModel := r.URL.Query().Get("model")
	if dMake == "" {
		http.Error(w, "Missing device make query", http.StatusBadRequest)
		return
	}
	if dModel == "" {
		http.Error(w, "Missing device model query", http.StatusBadRequest)
		return
	}
	err := db.DeleteDeviceModel(conn, dMake, dModel)
	if err != nil {
		http.Error(w, "Error deleting device model", http.StatusInternalServerError)
		return
	}
}

func DeviceModelHandler(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		handleGetDeviceModels(w, r, conn)
	case http.MethodPost:
		postDeviceModel(w, r, conn)
	case http.MethodPut:
		putDeviceModel(w, r, conn)
	case http.MethodPatch:
		patchDeviceModel(w, r, conn)
	case http.MethodDelete:
		deleteDeviceModel(w, r, conn)
	}
}
