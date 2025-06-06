package query

import (
	"context"
	"fmt"

	"gw-exchanger/domain/repository"
)

type UserRepository struct {
	client repository.Client
}

func NewRepository(client repository.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) GettingCourse(ctx context.Context) (map[string]float64, error) {
	courses := make(map[string]float64)
	var usdRub, usdEur, eurRub float64

	err := r.client.QueryRow(ctx, `SELECT * FROM exchanger `).Scan(
		&usdRub, &usdEur, &eurRub)
	if err != nil {
		return nil, fmt.Errorf("failed to get сourse: %w", err)
	}

	courses["USD_RUB"] = usdRub / 10000
	courses["USD_EUR"] = usdEur / 10000
	courses["EUR_RUB"] = eurRub / 10000

	return courses, nil
}

func (r *UserRepository) Exchange(ctx context.Context, fromCurrency, toCurrency string) (int, error) {
	var rate int

	err := r.client.QueryRow(ctx, `
        SELECT
            CASE
                WHEN $1 = 'USD' AND $2 = 'RUB' THEN usd_rub
                WHEN $1 = 'USD' AND $2 = 'EUR' THEN usd_eur
                WHEN $1 = 'EUR' AND $2 = 'RUB' THEN eur_rub
                WHEN $1 = 'EUR' AND $2 = 'USD' THEN eur_usd
                WHEN $1 = 'RUB' AND $2 = 'EUR' THEN rub_eur
                WHEN $1 = 'RUB' AND $2 = 'USD' THEN rub_usd
                ELSE 0
            END
        FROM exchanger
        LIMIT 1
    `, fromCurrency, toCurrency).Scan(&rate)

	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	return rate, nil
}
