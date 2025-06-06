package query

import (
	"context"
	"fmt"
	"gw-exchanger/domain/models"
	"gw-exchanger/domain/repository"
)

type UserRepository struct {
	client repository.Client
}

func NewRepository(client repository.Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) GettingCourse(ctx context.Context) (*models.Сourse, error) {
	var сourse models.Сourse
	err := r.client.QueryRow(ctx, `SELECT * FROM exchanger `).Scan(
		&сourse.USD, &сourse.RUB, &сourse.EUR)
	if err != nil {
		return nil, fmt.Errorf("failed to get сourse: %w", err)
	}

	return &сourse, nil
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
