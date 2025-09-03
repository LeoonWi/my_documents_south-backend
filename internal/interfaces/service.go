package interfaces

import (
	"context"
	"my_documents_south_backend/internal/model"
)

type EmployeeService interface{}

type EmployeeSpecsService interface{}

type RequestService interface{}

type RoleService interface {
	CreateRole(ctx context.Context, name string) (*model.Role, error)
	GetRoles(ctx context.Context) *[]model.Role
	GetRoleById(ctx context.Context, id int) (*model.Role, error)
}

type ServiceService interface {
	CreateService(ctx context.Context, name string) (*model.Service, error)
	GetService(ctx context.Context) *[]model.Service
	GetServiceById(ctx context.Context, id int) (*model.Service, error)
}

type TariffService interface {
	CreateTariff(ctx context.Context, name string) (*model.Tariff, error)
	GetTariffs(ctx context.Context) *[]model.Tariff
	GetTariffByID(ctx context.Context, id int) (*model.Tariff, error)
	GetTariffByName(ctx context.Context, name string) (*model.Tariff, error)
	UpdateTariff(ctx context.Context, id int, name string) (*model.Tariff, error)
	DeleteTariff(ctx context.Context, id int) error
}

type UserService interface{}
