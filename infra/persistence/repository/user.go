package repository

import (
	"context"

	"github.com/inari111/money-transfer/infra/persistence/mysql"

	"github.com/go-gorp/gorp"

	"github.com/inari111/money-transfer/domain/user"
)

func NewUserRepository(dbmap *gorp.DbMap) user.Repository {
	dbmap.AddTableWithName(mysql.User{}, "user").SetKeys(true, "id")
	return &userRepository{
		dbmap: dbmap,
	}
}

type userRepository struct {
	dbmap *gorp.DbMap
}

func (*userRepository) Get(ctx context.Context, id user.ID) (*user.User, error) {
	panic("implement me")
}

func (r *userRepository) Put(ctx context.Context, user *user.User) error {
	defer r.dbmap.Db.Close()

	u := &mysql.User{
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return r.dbmap.Insert(u)
}
