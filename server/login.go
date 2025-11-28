package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"smart_house/be/db"
	"smart_house/be/server/requests"
	"smart_house/be/server/responses"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var creds requests.Login
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Login Invalid Request Body", http.StatusBadRequest)
		return
	}

	// get the user
	// check the password
	user, err := db.ReadUser(conn, creds.Username)
	if err != nil {
		http.Error(w, "Login Invalid Username", http.StatusUnauthorized)
		return
	}

	// will need to update this with hashing and de-hashing if fe ends up using hashing

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Login Invalid Password", http.StatusUnauthorized)
		return
	}

	session, _ := db.ReadSession(conn, creds.Username)
	if session == nil {
		log.Print("DB Read Session Not Found")
		err := db.CreateSession(conn, creds.Username)
		if err != nil {
			http.Error(w, "Login Invalid Username", http.StatusUnauthorized)
		}
		session, _ = db.ReadSession(conn, creds.Username)
	} else {
		err := db.UpdateSession(conn, creds.Username)
		if err != nil {
			http.Error(w, "Login Invalid Username", http.StatusUnauthorized)
		}
		session, _ = db.ReadSession(conn, creds.Username)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.Success{Success: true})
}
