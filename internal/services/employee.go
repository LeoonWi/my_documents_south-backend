package services

import (
	"context"
	"my_documents_south_backend/internal/models"
	"time"
)

type employeeService struct {
	employeeRepository models.EmployeeRepository
	contextTimeout     time.Duration
}

func NewEmployeeService(employeeRepository models.EmployeeRepository, contextTimeout time.Duration) models.EmployeeService {
	return &employeeService{employeeRepository: employeeRepository, contextTimeout: contextTimeout}
}

func (s *employeeService) Create(c context.Context, name string) (*models.Employee, error) {
	// TODO create employee service
	return nil, nil
}

func (s *employeeService) Get(c context.Context) *[]models.Employee {
	// TODO get list employee service
	return nil
}

func (s *employeeService) GetById(c context.Context, id int) (*models.Employee, error) {
	// TODO get by id employee user service
	return nil, nil
}

func (s *employeeService) Update(c context.Context, id int, name string) (*models.Employee, error) {
	// TODO update employee service
	// DONT TOUCH
	return nil, nil
}

func (s *employeeService) Delete(c context.Context, id int) error {
	// TODO delete employee service
	return nil
}
