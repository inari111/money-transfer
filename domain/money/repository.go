package money

import "context"

type Repository interface {
	Get(ctx context.Context, id ID) (*Money, error)
	Put(ctx context.Context, money *Money) error
}
