package model

import (
	"database/sql"
	"time"
)

type User struct {
	UUID      string       `db:"uuid"`
	Profile   UserProfile  `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserProfile struct {
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Age       int64        `db:"age"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
