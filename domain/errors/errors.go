package errors

import "github.com/pkg/errors"

type ErrorType int

const (
	Unknown ErrorType = 0

	BadRequest   = 400
	Unauthorized = 401
	Forbidden    = 403
	NotFound     = 404

	InternalError = 500
)

// TODO: StackTrace必要だったら追加する
// struct名はErrorの方がいいかもしれない
type Fundamental struct {
	error
	ErrorType
}

func (f Fundamental) ErrType() ErrorType {
	return f.ErrorType
}

func New(message string, errType ErrorType) error {
	err := errors.New(message)
	return Fundamental{
		error:     err,
		ErrorType: errType,
	}
}

func ErrTypeFrom(err error) ErrorType {
	domainErr, ok := err.(Fundamental)
	if !ok {
		return Unknown
	}
	return domainErr.ErrType()
}
