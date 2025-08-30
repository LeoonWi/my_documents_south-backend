package employee_specs

import (
	"my_documents_south_backend/internal/domain/employee"
	"my_documents_south_backend/internal/domain/service"
	"time"
)

type Model struct {
	Id int64 `json:"id,omitempty" db:"id"`

	EmployeeId int64          `json:"employee_id,omitempty" db:"employee_id"`
	Employee   employee.Model `json:"employee,omitempty" db:"-"`

	ServiceId int           `json:"service_id,omitempty" db:"service_id"`
	Service   service.Model `json:"service,omitempty" db:"-" `

	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
