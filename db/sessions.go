package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func CreateSession(conn *sql.DB, username string) {
	token := uuid.New().String()
	created := time.Now().UTC().String()
	expires := time.Now().UTC().Add(24 * time.Hour).String()
	cmd := fmt.Sprintf("INSERT INTO sessions (username, token, created, expiry) VALUES ('%s', '%s', '%s', '%s');", username, token, created, expires)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Create Session: ", err)
	}
}

func UpdateSession(conn *sql.DB, username string) {
	//token := uuid.New().String()
	//created := time.Now().UTC().String()
	//expires := time.Now().UTC().Add(24 * time.Hour).String()
	// update token, created time, expires time
}
