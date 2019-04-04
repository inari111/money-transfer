package repository

import (
	"context"

	"github.com/inari111/money-transfer/domain"

	"github.com/go-gorp/gorp"
	"github.com/inari111/money-transfer/infra/persistence/mysql"

	"github.com/inari111/money-transfer/domain/user"
)

func NewUserProfileRepository(dbmap *gorp.DbMap, now domain.CurrentTimeFunc) user.ProfileRepository {
	dbmap.AddTableWithName(mysql.User{}, "user").SetKeys(true, "id")
	return &userProfileRepository{
		dbmap: dbmap,
		now:   now,
	}
}

type userProfileRepository struct {
	dbmap *gorp.DbMap
	now   domain.CurrentTimeFunc
}

func (*userProfileRepository) Get(ctx context.Context, id user.ID) (*user.Profile, error) {
	panic("implement me")
}

func (r *userProfileRepository) Put(ctx context.Context, profile *user.Profile) error {
	tx, err := r.dbmap.Begin()
	if err != nil {
		return err
	}

	var u mysql.User
	err = tx.SelectOne(&u, "select * from user where id=?", profile.ID)
	if err != nil {
		return err
	}

	_, err = tx.Update(u.WithProfile(profile, r.now()))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
