package models

import (
	"time"
)

type Session struct {
	Id        uint      `db:"id" json:"id"`
	UserId  []uint8   `db:"owner_id" json:"owner_id"`
	Stale        bool      `db:"stale" json:"stale"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	ExpiredAt time.Time `db:"expired_at" json:"expired_at"`
}
