package model

import (
	"time"
)

type Role struct {
	Id        int       `json:"id,omitempty" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
