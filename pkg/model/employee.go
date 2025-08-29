package model

import (
	"time"
)

type Employee struct {
	Id         int    `json:"id,omitempty" db:"id"`
	Name       string `json:"name,omitempty" db:"name"`
	LastName   string `json:"last_name,omitempty" db:"last_name"`
	MiddleName string `json:"middle_name,omitempty" db:"middle_name"`
	Email      string `json:"email,omitempty" db:"email"`

	RoleId int  `json:"role_id,omitempty" db:"role_id"`
	Role   Role `json:"role,omitempty" db:"-"`

	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
