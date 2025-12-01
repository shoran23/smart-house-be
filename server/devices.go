package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smart_house/be/db"
	"smart_house/be/db/models"
)

func getDevices(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	dm, err := db.ReadAllDevices(conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getDevice(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing device name", http.StatusBadRequest)
		return
	}
	d, err := db.ReadDevice(conn, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetDevices(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	params := r.URL.Query()
	if len(params) > 0 {
		getDevice(w, r, conn)
	} else {
		getDevices(w, r, conn)
	}
}

func postDevice(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var d models.Device
	if r.Body == nil {
		http.Error(w, "Missing device body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.CreateDevice(conn, d.Name, d.Room, d.Location, d.DeviceId, d.Address, d.Make, d.Model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func putDevice(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var d models.Device
	if r.Body == nil {
		http.Error(w, "Missing device body", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.CreateDevice(conn, d.Name, d.Room, d.Location, d.DeviceId, d.Address, d.Make, d.Model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func patchDevice(w http.ResponseWriter, r *http.Request, conn *sql.DB) {

}

func deleteDevice(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing device name", http.StatusBadRequest)
		return
	}
	err := db.DeleteDevice(conn, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeviceHandler(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		handleGetDevices(w, r, conn)
	case http.MethodPost:
		postDevice(w, r, conn)
	case http.MethodPut:
		putDevice(w, r, conn)
	case http.MethodPatch:
		patchDevice(w, r, conn)
	case http.MethodDelete:
		deleteDevice(w, r, conn)
	}
}
