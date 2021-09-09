package models

import (
	"time"
)

type Org struct {
	Id        []uint8 `db:"id" json:"id"`
	Sub       string  `db:"sub" json:"sub"`
	Name  string  `db:"name" json:"name"`
	OwnerUserId []uint8 `db:"owner_user_id" json:"owner_user_id"`
	OwnerUser User
	Users  []User
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
