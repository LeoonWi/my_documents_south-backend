package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dongri/phonenumber"
	"my_documents_south_backend/internal/models"
	"regexp"
	"time"
	"unicode"
)

type userService struct {
	tariffRepository models.TariffRepository
	userRepository   models.UserRepository
	contextTimeout   time.Duration
}

func NewUserService(userRepository models.UserRepository, tariffRepository models.TariffRepository, contextTimeout time.Duration) models.UserService {
	return &userService{userRepository: userRepository, tariffRepository: tariffRepository, contextTimeout: contextTimeout}
}

func (s *userService) Create(c context.Context, user *models.User, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	//я так понял,у нас будет валидация всех полей и на фронте, и здесь. может я ничего не смыслю,но зачем?
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return nil, errors.New("invalid email format")
	}

	normalized := phonenumber.Parse(user.Phone, "RU")
	if normalized == "" {
		return nil, errors.New("invalid phone number")
	}
	user.Phone = normalized

	// я так подумал,у нас же на этом моменте будет пароль зашифрован и пересолен,валидность тогда нужна на фронте, а не здесь(пока просто оставлю,как есть)
	//ну или после валидации сразу шифровать пароль.я бы это ,кстати, сделал,если бы мог увидеть свой результат.
	//а как я увижу результат без колонки "password" в бд?
	if len(password) < 8 {
		return nil, errors.New("invalid password: must contain at least 8 characters")
	}
	hasLetter := false
	hasDigit := false
	for _, ch := range password {
		if unicode.IsLetter(ch) {
			hasLetter = true
		}
		if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return nil, errors.New("invalid password: must contain at least one letter and one digit")
	}

	id := 1
	tariff := &models.Tariff{Id: id}
	terr := s.tariffRepository.GetById(ctx, id, tariff)
	if terr != nil {
		return nil, fmt.Errorf("failed to check default tariff: %w", terr)
	}

	// тариф по умолчанию (id = 1)
	user.TariffId = id

	err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *userService) Get(c context.Context) *[]models.User {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	var user []models.User
	err := s.userRepository.Get(ctx, &user)
	if err != nil {
		return nil
	}
	return &user
}

func (s *userService) GetById(c context.Context, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if id < 1 {
		return nil, errors.New("invalid id")
	}

	user := &models.User{Id: int64(id)}
	err := s.userRepository.GetById(ctx, id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Update(c context.Context, id int, name string) (*models.User, error) {
	// TODO update user service
	// DONT TOUCH
	return nil, nil
}

func (s *userService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.userRepository.Delete(ctx, id)
}
