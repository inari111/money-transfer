package money

import (
	"time"

	"github.com/inari111/money-transfer-study/domain/user"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Money struct {
	ID        ID
	UserID    user.ID
	Amount    uint64
	UpdatedAt time.Time
}
