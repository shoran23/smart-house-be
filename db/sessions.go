package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"
	"time"

	"github.com/google/uuid"
)

func CreateSession(conn *sql.DB, username string) error {
	token := uuid.New().String()
	created := time.Now().UTC().String()
	expires := time.Now().UTC().Add(24 * time.Hour).String()
	cmd := fmt.Sprintf("INSERT INTO sessions (username, token, created, expiry) VALUES ('%s', '%s', '%s', '%s');", username, token, created, expires)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Print("DB Create Session: ", err)
		return err
	}
	return nil
}

func ReadSession(conn *sql.DB, username string) (*models.Session, error) {
	cmd := fmt.Sprintf("SELECT * FROM sessions WHERE username = '%s';", username)
	row := conn.QueryRow(cmd)

	var session models.Session
	err := row.Scan(&session.Token, &session.Username, &session.Created, &session.Expiry)
	if err != nil {
		log.Print("DB Read Session: ", err)
		return nil, err
	}
	return &session, nil
}

func UpdateSession(conn *sql.DB, username string) error {
	token := uuid.New().String()
	created := time.Now().UTC().String()
	expires := time.Now().UTC().Add(24 * time.Hour).String()
	cmd := fmt.Sprintf("UPDATE sessions SET token = '%s', created = '%s', expiry = '%s' WHERE username = '%s';", token, created, expires, username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Update Session: ", err)
		return err
	}
	return nil
}

func DeleteSession(conn *sql.DB, username string) error {
	cmd := fmt.Sprintf("DELETE FROM sessions WHERE username = '%s';", username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Println("DB Delete Session: ", err)
		return err
	}
	return nil
}
