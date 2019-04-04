package repository

import (
	"context"

	"github.com/inari111/money-transfer/infra/persistence/mysql"

	"github.com/go-gorp/gorp"

	"github.com/inari111/money-transfer/domain/money"
)

func NewMoneyRepository(dbmap *gorp.DbMap) money.Repository {
	dbmap.AddTableWithName(mysql.Money{}, "money").SetKeys(true, "id")
	return &moneyRepository{
		dbmap: dbmap,
	}
}

type moneyRepository struct {
	dbmap *gorp.DbMap
}

func (*moneyRepository) Get(ctx context.Context, id money.ID) (*money.Money, error) {
	panic("implement me")
}

func (r *moneyRepository) Put(ctx context.Context, money *money.Money) error {
	m := &mysql.Money{
		UserID:    money.UserID.Int64(),
		Amount:    int64(money.Amount),
		UpdatedAt: money.UpdatedAt,
	}
	return r.dbmap.Insert(m)
}
