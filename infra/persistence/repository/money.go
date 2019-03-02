package repository

import (
	"context"

	"github.com/inari111/money-transfer/domain/money"
)

func NewMoneyRepository() money.Repository {
	return &moneyRepository{}
}

type moneyRepository struct{}

func (*moneyRepository) Get(ctx context.Context, id money.ID) (*money.Money, error) {
	panic("implement me")
}

func (*moneyRepository) Put(ctx context.Context, money *money.Money) error {
	panic("implement me")
}
