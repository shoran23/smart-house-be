package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(conn *sql.DB, username, password string, priv models.Privilege) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("DB Create User Crypt: ", err)
		return
	}
	cmd := fmt.Sprintf("INSERT INTO users (username, password, privilege) VALUES ('%s', '%s', %d);", username, string(bytes), priv)
	_, err = conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Create User Execute: ", err)
	}
}

func ReadAllUsers(conn *sql.DB, search string) *[]models.User {
	cmd := fmt.Sprintf("SELECT username, privilege FROM users;")
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Fatal("DB Read All Users: ", err)
	}
	defer rows.Close()

	var Users []models.User
	for rows.Next() {
		var User models.User
		err := rows.Scan(&User.Username, &User.Privilege)
		if err != nil {
			log.Fatal("DB Read All Users: ", err)
		}
		Users = append(Users, User)
	}
	return &Users
}

func ReadUser(conn *sql.DB, username string) *models.User {
	cmd := fmt.Sprintf("SELECT * FROM users WHERE username = '%s';", username)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Fatal("DB Read User Query: ", err)
		return nil
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			log.Fatal("DB Read User Scan: ", err)
		}
	}
	return &user
}

func UpdateUserUsername(conn *sql.DB, username, new string) {
	cmd := fmt.Sprintf("UPDATE users SET username = '%s' WHERE username = '%s';", new, username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Update User Username: ", err)
	}
}

func UpdateUserPassword(conn *sql.DB, username, new string) {
	cmd := fmt.Sprintf("UPDATE users SET password = '%s' WHERE username = '%s';", new, username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Update User Password Execute: ", err)
	}
}

func UpdateUserPrivilege(conn *sql.DB, username string, new models.Privilege) {
	cmd := fmt.Sprintf("UPDATE users SET privilege = %d WHERE username = '%s';", new, username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Update User Privilege: ", err)
	}
}

func DeleteUser(conn *sql.DB, username string) {
	cmd := fmt.Sprintf("DELETE FROM users WHERE username = '%s'", username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("DB Delete User Execute: ", err)
	}
}
