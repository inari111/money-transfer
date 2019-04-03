package mysql

import (
	"database/sql"

	"github.com/inari111/money-transfer/domain/errors"
)

func ToDomainError(err error) error {
	if err == sql.ErrNoRows {
		return errors.New("not found", errors.NotFound)
	}
	return errors.New(err.Error(), errors.InternalError)
}
