package mysql

import "time"

type Money struct {
	ID        int64     `db:"id, primarykey"`
	UserID    int64     `db:"user_id"`
	Amount    int64     `db:"amount"`
	UpdatedAt time.Time `db:"updated_at"`
}
