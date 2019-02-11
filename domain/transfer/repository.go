package transfer

import "context"

type Repository interface {
	Get(ctx context.Context, id ID) (*Transfer, error)
	Put(ctx context.Context, transfer *Transfer) error
}
