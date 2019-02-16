package repository

import (
	"context"

	"github.com/inari111/money-transfer-study/domain/user"
)

func NewUserProfileRepository() user.ProfileRepository {
	return &userProfileRepository{}
}

type userProfileRepository struct{}

func (*userProfileRepository) Get(ctx context.Context, id user.ID) (*user.Profile, error) {
	panic("implement me")
}

func (*userProfileRepository) Put(ctx context.Context, profile *user.Profile) error {
	panic("implement me")
}
