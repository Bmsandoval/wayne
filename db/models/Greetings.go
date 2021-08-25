package models

import "time"

type Greetings struct {
	Id int `db:"id" json:"greeting_id"`
	Value string `db:"value" json:"value"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

func (m Greetings) TableName() string {
	return "greetings"
}

func (m Greetings) ModelName() string {
	return "GreetingModel"
}