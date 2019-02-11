package repository

import (
	"context"

	"github.com/inari111/money-transfer-study/domain/user"
)

func NewProfileRepository() user.ProfileRepository {
	return &profileRepository{}
}

type profileRepository struct{}

func (*profileRepository) Get(ctx context.Context, id user.ID) (*user.Profile, error) {
	panic("implement me")
}

func (*profileRepository) Put(ctx context.Context, profile *user.Profile) error {
	panic("implement me")
}
