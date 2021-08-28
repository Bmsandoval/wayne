package models

import (
	"time"
)

type User struct {
	Id        []uint8 `db:"id" json:"id"`
	Sub       string  `db:"sub" json:"sub"`
	Username  string  `db:"username" json:"username"`
	Password  string  `db:"password" json:"password"`
	Sessions  []Session
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
