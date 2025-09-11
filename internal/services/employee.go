package services

import (
	"context"
	"errors"
	"fmt"
	"my_documents_south_backend/internal/middleware"
	"my_documents_south_backend/internal/models"
	"my_documents_south_backend/internal/utils/password"
	"regexp"
	"time"
	"unicode"
)

type employeeService struct {
	roleRepository     models.RoleRepository
	employeeRepository models.EmployeeRepository
	contextTimeout     time.Duration
}

func NewEmployeeService(employeeRepository models.EmployeeRepository, roleRepository models.RoleRepository, contextTimeout time.Duration) models.EmployeeService {
	return &employeeService{employeeRepository: employeeRepository, roleRepository: roleRepository, contextTimeout: contextTimeout}
}

func (s *employeeService) Create(c context.Context, employee *models.Employee) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	hasLetter := false
	hasDigit := false

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(employee.Email) {
		return errors.New("invalid email format")
	}

	if len(employee.Password) < 6 {
		return errors.New("invalid password: must contain at least 6 characters")
	}

	for _, ch := range employee.Password {
		if unicode.IsLetter(ch) {
			hasLetter = true
		}
		if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return errors.New("invalid password: must contain at least one letter and one digit")
	}

	var role models.Role
	err := s.roleRepository.GetById(ctx, employee.RoleId, &role)
	if err != nil {
		return fmt.Errorf("failed to check default role: %w", err)
	}

	employee.Password, err = password.Encrypt(employee.Password)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	err = s.employeeRepository.Create(ctx, employee)
	if err != nil {
		return err
	}
	return nil
}

func (s *employeeService) Login(c context.Context, employee *models.Employee) (string, string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	var employee2 models.Employee
	err := s.employeeRepository.GetByEmail(ctx, employee.Email, &employee2)
	if err != nil {
		return "", "", err
	}

	if err := password.Compare(employee2.Password, employee.Password); err != nil {
		return "", "", err
	}

	accessToken, err := middleware.JWTGenerate(employee2.Id, &employee2.RoleId, time.Hour)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := middleware.JWTGenerate(employee2.Id, &employee2.RoleId, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *employeeService) Get(c context.Context) *[]models.Employee {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	var employee []models.Employee
	err := s.employeeRepository.Get(ctx, &employee)
	if err != nil {
		return nil
	}
	return &employee
}

func (s *employeeService) GetById(c context.Context, id int) (*models.Employee, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if id < 1 {
		return nil, errors.New("invalid id")
	}

	employee := &models.Employee{Id: int64(id)}
	err := s.employeeRepository.GetById(ctx, id, employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *employeeService) Update(c context.Context, id int, employee *models.Employee) error {
	// TODO update employee service
	// DONT TOUCH
	return nil
}

func (s *employeeService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	return s.employeeRepository.Delete(ctx, id)
}
