package services

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository interface {
	GettingCourse(ctx context.Context) (map[string]float64, error)
	Exchange(ctx context.Context, fromCurrency, toCurrency string) (int, error)
}

type UserService struct {
	repo UserRepository
	log  *slog.Logger
}

func NewUserService(log *slog.Logger, repo UserRepository) *UserService {
	return &UserService{log: log, repo: repo}
}

func (s *UserService) GetExchange(ctx context.Context) (map[string]float64, error) {
	s.log.Info("receiving exchange")
	return s.repo.GettingCourse(ctx)
}

func (s *UserService) GetRate(ctx context.Context, fromCurrency, toCurrency string, amount int) (float64, error) {
	rate, err := s.repo.Exchange(ctx, fromCurrency, toCurrency)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "internal error")
	}

	result := (float64(amount) / 100) * float64(rate) / 1000000

	return float64(result), nil
}
