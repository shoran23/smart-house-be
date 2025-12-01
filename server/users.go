package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smart_house/be/db"
	"smart_house/be/db/models"
)

func getUsers(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	users, err := db.ReadAllUsers(conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	param := r.URL.Query().Get("username")
	if param == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}
	user, err := db.ReadUser(conn, param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
}

func handleGetUsers(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	params := r.URL.Query()
	if len(params) > 0 {
		getUser(w, r, conn)
	} else {
		getUsers(w, r, conn)
	}
}

func postUser(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var user models.User
	if r.Body == nil {
		http.Error(w, "Missing body", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.CreateUser(conn, user.Username, user.Password, models.Privilege(user.Privilege))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func putUser(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	param := r.URL.Query().Get("username")
	if param == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.ReplaceUser(conn, user.Username, user.Password, models.Privilege(user.Privilege))
}

func patchUser(w http.ResponseWriter, r *http.Request, conn *sql.DB) {

}

func deleteUser(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	param := r.URL.Query().Get("username")
	if param == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}
	err := db.DeleteUser(conn, param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r, conn)
	case http.MethodPost:
		postUser(w, r, conn)
	case http.MethodPut:
		putUser(w, r, conn)
	case http.MethodPatch:
		patchUser(w, r, conn)
	case http.MethodDelete:
		deleteUser(w, r, conn)
	}
}
