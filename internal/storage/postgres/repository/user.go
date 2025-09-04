package repository

import (
	"context"
	"my_documents_south_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) models.UserRepository {
	return &userRepository{conn: conn}
}

func (r *userRepository) Create(c context.Context, user *models.User) error {
	// TODO create user repository
	return nil
}

func (r *userRepository) Get(c context.Context, user *[]models.User) error {
	// TODO get user repository
	return nil
}

func (r *userRepository) GetById(c context.Context, id int, user *models.User) error {
	// TODO get by id user repository
	return nil
}

func (r *userRepository) Update(c context.Context, user *models.User) error {
	// TODO update user repository
	// DONT TOUCH
	return nil
}

func (r *userRepository) Delete(c context.Context, id int) error {
	// TODO delete user repository
	return nil
}
