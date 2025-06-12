package services

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
)

type UserRepository interface {
	GettingCourse(ctx context.Context) (map[string]float64, error)
	Exchange(ctx context.Context, fromCurrency, toCurrency string) (int, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetExchange(ctx context.Context) (map[string]float64, error) {
	return s.repo.GettingCourse(ctx)
}

func (s *UserService) GetRate(ctx context.Context, fromCurrency, toCurrency string, amount int) (float64, error) {
	rate, err := s.repo.Exchange(ctx, fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}

	result := (float64(amount) / 100) * float64(rate) / 1000000
	log.Info(123)
	return float64(result), nil
}
