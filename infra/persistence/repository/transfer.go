package repository

import (
	"context"

	"github.com/inari111/money-transfer-study/domain/transfer"
)

func NewTransferRepository() transfer.Repository {
	return &transferRepository{}
}

type transferRepository struct{}

func (*transferRepository) Get(ctx context.Context, id transfer.ID) (*transfer.Transfer, error) {
	panic("implement me")
}

func (*transferRepository) Put(ctx context.Context, transfer *transfer.Transfer) error {
	panic("implement me")
}
