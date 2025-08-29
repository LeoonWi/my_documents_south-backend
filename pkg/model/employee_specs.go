package model

import (
	"time"
)

type Employee_specs struct {
	Id         int      `json:"id,omitempty" db:"id"`
	EmployeeId int      `json:"employee_id,omitempty" db:"employee_id"`
	Employee   Employee `json:"employee,omitempty" db:"-"`

	ServiceId int     `json:"service_id,omitempty" db:"service_id"`
	Service   Service `json:"service,omitempty" db:"-" `

	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
