package mysql

import (
	"time"

	"github.com/inari111/money-transfer/domain/user"
)

type User struct {
	ID int64 `db:"id, primarykey, autoincrement"`

	Name string `db:"name"`
	Age  int    `db:"age"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) WithProfile(profile *user.Profile, now time.Time) *User {
	user := *u

	user.Name = profile.Name
	user.Age = profile.Age
	user.UpdatedAt = now

	return &user
}

func (u *User) ToDomain() *user.User {
	return &user.User{
		ID:        user.ID(u.ID),
		Profile:   u.ToDomainProfile(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToDomainProfile() *user.Profile {
	return &user.Profile{
		ID:   user.ID(u.ID),
		Name: u.Name,
		Age:  u.Age,
	}
}
