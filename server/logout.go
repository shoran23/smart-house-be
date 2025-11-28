package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"smart_house/be/server/requests"
	"smart_house/be/server/responses"

	"smart_house/be/db"
)

func Logout(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Logout Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var username requests.Logout
	if err := json.NewDecoder(r.Body).Decode(&username); err != nil {
		http.Error(w, "Logout Invalid Request Body", http.StatusBadRequest)
		return
	}

	session, err := db.ReadSession(conn, username.Username)
	if err != nil {
		http.Error(w, "Logout Invalid Username", http.StatusUnauthorized)
		return
	}

	if session == nil {
		http.Error(w, "Logout Invalid Username", http.StatusUnauthorized)
		return
	} else {
		fmt.Printf("Logout Delete Session %s\n", session.Username)
		err = db.DeleteSession(conn, session.Username)
		if err != nil {
			http.Error(w, "Logout Invalid Username", http.StatusUnauthorized)
			return
		}
		err := db.DeleteSession(conn, session.Username)
		if err != nil {
			http.Error(w, "Logout Invalid Username", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses.Success{Success: true})
	}
}
