package money

import (
	"context"

	"github.com/inari111/money-transfer/domain"

	"github.com/inari111/money-transfer/domain/money"
)

type Application interface {
	Gain(
		ctx context.Context,
		money *money.Money,
	) error
}

func NewApplication(
	moneyRepo money.Repository,
	now domain.CurrentTimeFunc,
) Application {
	return &moneyApp{
		moneyRepo: moneyRepo,
		now:       now,
	}
}

type moneyApp struct {
	moneyRepo money.Repository
	now       domain.CurrentTimeFunc
}

func (a *moneyApp) Gain(
	ctx context.Context,
	money *money.Money,
) error {
	money.UpdatedAt = a.now()
	return a.moneyRepo.Put(ctx, money)
}
