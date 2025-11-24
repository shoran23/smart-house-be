package models

import "time"

type Session struct {
	Token    string
	Username string
	Created  time.Time
	Expiry   time.Time
}
