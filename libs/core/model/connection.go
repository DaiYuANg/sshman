package model

import (
	"time"
)

type SSHConnection struct {
	ID          string
	Name        string
	Host        string
	Port        int
	Username    string
	AuthMethod  string
	Password    string
	PrivateKey  []byte
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
