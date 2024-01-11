package model

import (
	"database/sql"
	"time"
)

type User struct {
	UUID      string
	Profile   UserProfile
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserProfile struct {
	FirstName string
	LastName  string
	Age       int64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
