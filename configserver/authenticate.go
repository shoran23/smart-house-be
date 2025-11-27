package configserver

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smart_house/be/configserver/models"
	"smart_house/be/db"
)

func authenticate(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var creds models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Authenticate Invalid Request Body", http.StatusBadRequest)
		return
	}

	// get the user
	// check the password
	user, err := db.ReadUser(conn, creds.Username)
	if err != nil {
		http.Error(w, "Authenticate Invalid Username", http.StatusUnauthorized)
		return
	}

	// will need to update this with hashing and de-hashing if fe ends up using hashing
	if user.Password != creds.Password {
		http.Error(w, "Authenticate Invalid Password", http.StatusUnauthorized)
		return
	}

	// create or update session
	// check session exists
	// if it does then update it
	// if not create one
	// add session token and username to cookies
	// create response
	// add cookies to response
	// add body to response
	// send response
}
