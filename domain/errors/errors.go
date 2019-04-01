package errors

import "github.com/pkg/errors"

type ErrorType int

const (
	BadRequest   ErrorType = 400
	Unauthorized ErrorType = 401
	Forbidden    ErrorType = 403
	NotFound     ErrorType = 404

	InternalError ErrorType = 500
)

// TODO: StackTrace必要だったら追加する
// struct名はErrorの方がいいかもしれない
type Fundamental struct {
	error
	ErrorType
}

func New(message string, errType ErrorType) error {
	err := errors.New(message)
	return Fundamental{
		error:     err,
		ErrorType: errType,
	}
}
