package services

import (
	"context"
	"my_documents_south_backend/internal/models"
	"time"
)

type userService struct {
	userRepository models.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepository models.UserRepository, contextTimeout time.Duration) models.UserService {
	return &userService{userRepository: userRepository, contextTimeout: contextTimeout}
}

func (s *userService) Create(c context.Context, name string) (*models.User, error) {
	// TODO create user service
	return nil, nil
}

func (s *userService) Get(c context.Context) *[]models.User {
	// TODO get user service
	return nil
}

func (s *userService) GetById(c context.Context, id int) (*models.User, error) {
	// TODO get by id user service
	return nil, nil
}

func (s *userService) Update(c context.Context, id int, name string) (*models.User, error) {
	// TODO update user service
	// DONT TOUCH
	return nil, nil
}

func (s *userService) Delete(c context.Context, id int) error {
	// TODO delete user service
	return nil
}
