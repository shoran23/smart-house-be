package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "./db/smart_house.sql")
	if err != nil {
		return nil, err
	}
	setup(conn)
	return conn, nil
}

func setup(conn *sql.DB) {
	path := "./db/sql/"
	files := []string{
		"create_users_table.sql",
		"create_sessions_table.sql",
		"create_device_models_table.sql",
		"create_devices_table.sql",
		//"populate_device_models_table.sql",
	}
	for _, file := range files {
		content, err := os.ReadFile(path + file)
		if err != nil {
			log.Fatalf("DB Setup Read File %s: %s", file, err)
			return
		}
		_, err = conn.Exec(string(content))
		if err != nil {
			log.Fatalf("DB Setup Execute File %s: %s", file, err)
			return
		}
	}
}
