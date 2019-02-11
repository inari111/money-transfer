package mysql

import "time"

type Transfer struct {
	ID int64 `db:"id, primarykey"`

	SrcUserID  string `db:"src_user_id"`
	DestUserID string `db:"dest_user_id"`

	Amount uint64 `db:"amount"`

	CreatedAt time.Time `db:"created_at"`
}
