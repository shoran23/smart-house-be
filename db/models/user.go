package models

type User struct {
	Username  string
	Password  string
	Privilege int
}

type Privilege int

const (
	AdminAccess Privilege = iota
	UserAccess
)
