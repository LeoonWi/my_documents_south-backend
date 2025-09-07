package models

import (
	"my_documents_south_backend/internal/interfaces"
	"time"
)

type Employee struct {
	Id         int64  `json:"id,omitempty" db:"id"`
	Name       string `json:"name,omitempty" db:"name"`
	LastName   string `json:"last_name,omitempty" db:"last_name"`
	MiddleName string `json:"middle_name,omitempty" db:"middle_name"`
	Email      string `json:"email,omitempty" db:"email"`
	Password   string `json:"password,omitempty" db:"password"`

	RoleId int  `json:"role_id,omitempty" db:"role_id"`
	Role   Role `json:"role,omitempty" db:"-"`

	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type EmployeeRepository interface {
	interfaces.EntityRepository[Employee]
}

type EmployeeService interface {
	interfaces.EntityService[Employee]
}
