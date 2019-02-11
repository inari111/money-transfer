package mysql

import "time"

type User struct {
	ID int64 `db:"id, primarykey, autoincrement"`

	Name string `db:"name"`
	Age  int    `db:"age"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
