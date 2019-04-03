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

func (r *userRepository) Get(ctx context.Context, id user.ID) (*user.User, error) {
	var u mysql.User
	err := r.dbmap.SelectOne(&u, "select * from user where id=?", id.String())
	if err != nil {
		return nil, mysql.ToDomainError(err)
	}
	return u.ToDomain(), nil
}

func (r *userRepository) Put(ctx context.Context, user *user.User) error {
	u := &mysql.User{
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return r.dbmap.Insert(u)
}
