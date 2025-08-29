package model

import (
	"time"
)

type Request struct {
	Id   int    `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`

	ServiceId  int      `json:"service_id,omitempty" db:"service_id"`
	Service    Service  `json:"service,omitempty" db:"-"`
	OwnerId    int      `json:"owner_id,omitempty" db:"owner_id"`
	User       User     `json:"user,omitempty" db:"-"`
	EmployeeId int      `json:"employee_id,omitempty" db:"employee_id"`
	Employee   Employee `json:"employee" db:"-"`

	Priority  int8      `json:"priority,omitempty" db:"priority"`
	Desc      string    `json:"desc,omitempty" db:"desc"`
	Status    int8      `json:"status,omitempty" db:"status"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DesiredAt time.Time `json:"desired_at,omitempty" db:"desired_at"`
	ClosedAt  time.Time `json:"closed_at,omitempty" db:"closed_at"`
}
