package services

import (
	"context"
	"errors"
	"my_documents_south_backend/internal/models"
	"time"
)

type requestService struct {
	requestRepository models.RequestRepository
	contextTimeout    time.Duration
}

func NewRequestService(requestRepository models.RequestRepository, contextTimeout time.Duration) models.RequestService {
	return &requestService{requestRepository: requestRepository, contextTimeout: contextTimeout}
}

func (s *requestService) Create(c context.Context, req *models.Request) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	err := s.requestRepository.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *requestService) Get(c context.Context) *[]models.Request {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	var req []models.Request
	err := s.requestRepository.Get(ctx, &req)
	if err != nil {
		return nil
	}
	return &req
}

func (s *requestService) GetById(c context.Context, id int) (*models.Request, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	if id < 1 {
		return nil, errors.New("invalid id")
	}

	req := &models.Request{Id: int64(id)}
	err := s.requestRepository.GetById(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *requestService) GetWithFilter(c context.Context, filter models.Request) ([]models.Request, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	var req []models.Request
	err := s.requestRepository.GetWithFilter(ctx, &req, filter)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (s *requestService) Update(c context.Context, id int, req *models.Request) error { return nil }

func (s *requestService) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	return s.requestRepository.Delete(ctx, id)
}
