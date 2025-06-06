package handlers

import (
	"context"
	"gw-exchanger/domain/models"
)

type UserService interface {
	GetExchange(ctx context.Context) (*models.Сourse, error)
	GetRate(ctx context.Context, fromCurrency, toCurrency string, amount int) (int, error)
}

func HandlerGetExchange(s UserService) {

}

func HandlerGetExchangeRateForCurrency(s UserService) (float64, error) {
	var fromCurrency, toCurrency string
	var amount int
	
	rate, err := s.GetRate(context.Background(), fromCurrency, toCurrency, amount)
	if err != nil {
		return 0, err
	}

	return float64(rate) / 10000, nil
}
