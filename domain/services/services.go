package services

import (
	"context"
	"gw-exchanger/domain/models"
)

type UserRepository interface {
	GettingCourse(ctx context.Context) (*models.Сourse, error)
	Exchange(ctx context.Context, fromCurrency, toCurrency string) (int, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetExchange(ctx context.Context) (*models.Сourse, error) {
	return s.repo.GettingCourse(ctx)
}

func (s *UserService) GetRate(ctx context.Context, fromCurrency, toCurrency string, amount int) (int, error) {
	rate, err := s.repo.Exchange(ctx, fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}
	result := rate * amount

	return result, nil
}
