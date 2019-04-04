package user

import (
	"time"
)

type ID int64

func (id ID) Int64() int64 {
	return int64(id)
}

type User struct {
	ID ID
	*Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}
