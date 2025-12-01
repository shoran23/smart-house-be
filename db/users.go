package db

import (
	"database/sql"
	"fmt"
	"log"
	"smart_house/be/db/models"

	"golang.org/x/crypto/bcrypt"
)

// generate password, add salt?
func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("DB Generate Password: ", err)
		return "", err
	}
	return string(bytes), nil
}

func CreateUser(conn *sql.DB, username, password string, priv models.Privilege) error {
	pass, err := generatePasswordHash(password)
	if err != nil {
		log.Print("DB Create User Crypt: ", err)
		return err
	}
	cmd := fmt.Sprintf("INSERT INTO users (username, password, privilege) VALUES ('%s', '%s', %d);", username, pass, priv)
	_, err = conn.Exec(cmd)
	if err != nil {
		log.Print("DB Create User Execute: ", err)
		return err
	}
	return nil
}

func ReadAllUsers(conn *sql.DB) (*[]models.User, error) {
	cmd := fmt.Sprintf("SELECT username, privilege FROM users;")
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Print("DB Read All Users: ", err)
		return nil, err
	}
	defer rows.Close()

	var Users []models.User
	for rows.Next() {
		var User models.User
		err := rows.Scan(&User.Username, &User.Privilege)
		if err != nil {
			log.Print("DB Read All Users: ", err)
			return nil, err
		}
		Users = append(Users, User)
	}
	return &Users, nil
}

func ReadUser(conn *sql.DB, username string) (*models.User, error) {
	cmd := fmt.Sprintf("SELECT * FROM users WHERE username = '%s';", username)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Print("DB Read User Query: ", err)
		return nil, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password, &user.Privilege)
		if err != nil {
			log.Print("DB Read User Scan: ", err)
			return nil, err
		}
	}
	return &user, nil
}

func UpdateUserUsername(conn *sql.DB, username, new string) error {
	cmd := fmt.Sprintf("UPDATE users SET username = '%s' WHERE username = '%s';", new, username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Print("DB Update User Username: ", err)
		return err
	}
	return nil
}

func UpdateUserPassword(conn *sql.DB, username, password string) error {
	pass, err := generatePasswordHash(password)
	if err != nil {
		log.Print("DB Update User Password: ", err)
		return err
	}
	cmd := fmt.Sprintf("UPDATE users SET password = '%s' WHERE username = '%s';", pass, username)
	_, err = conn.Exec(cmd)
	if err != nil {
		log.Print("DB Update User Password Execute: ", err)
		return err
	}
	return nil
}

func UpdateUserPrivilege(conn *sql.DB, username string, new models.Privilege) error {
	cmd := fmt.Sprintf("UPDATE users SET privilege = %d WHERE username = '%s';", new, username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Print("DB Update User Privilege: ", err)
		return err
	}
	return nil
}

func ReplaceUser(conn *sql.DB, username, password string, priv models.Privilege) error {
	pass, err := generatePasswordHash(password)
	cmd := fmt.Sprintf("UPDATE users SET username = '%s', password = '%s', privilige = '%d'  WHERE username = '%s';", username, pass, priv)
	_, err = conn.Exec(cmd)
	if err != nil {
		log.Print("DB ReplaceUser Execute: ", err)
		return err
	}
	return nil
}

func DeleteUser(conn *sql.DB, username string) error {
	cmd := fmt.Sprintf("DELETE FROM users WHERE username = '%s'", username)
	_, err := conn.Exec(cmd)
	if err != nil {
		log.Print("DB Delete User Execute: ", err)
		return err
	}
	return nil
}

func CreateDefaultUser(conn *sql.DB) error {
	user, err := ReadUser(conn, "admin")
	if err != nil {
		log.Print("DB Create Default User: ", err)
		return err
	}
	if user == nil {
		pass, err := generatePasswordHash("admin")
		if err != nil {
			log.Print("DB Create Default User: ", err)
			return err
		}
		err = CreateUser(conn, "admin", pass, models.AdminAccess)
		if err != nil {
			log.Print("DB Create Default User: ", err)
			return err
		}
	}

	return nil
}
