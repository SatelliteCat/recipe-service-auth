package model

import (
	"time"
)

type User struct {
	UUID      string
	Profile   UserProfile
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserProfile struct {
	FirstName string
	LastName  string
	Age       int64
	CreatedAt time.Time
	UpdatedAt *time.Time
}
