package transfer

import (
	"time"

	"github.com/inari111/money-transfer-study/domain/user"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Transfer struct {
	ID ID

	// 送金元
	Src user.ID
	// 送金先
	Dest user.ID

	Amount uint64

	CreatedAt time.Time
}
