package rpc

import (
	"github.com/inari111/money-transfer/domain/errors"
	"github.com/twitchtv/twirp"
)

var errCodeMap = map[errors.ErrorType]twirp.ErrorCode{
	errors.BadRequest:    twirp.Internal,
	errors.Unauthorized:  twirp.InvalidArgument,
	errors.Forbidden:     twirp.Unauthenticated,
	errors.NotFound:      twirp.NotFound,
	errors.InternalError: twirp.Internal,
}

func TwirpErrFrom(err error) error {
	f, ok := err.(errors.Fundamental)
	if !ok {
		return twirp.NewError(twirp.Internal, err.Error())
	}

	code, ok := errCodeMap[f.ErrorType]
	if !ok {
		return twirp.NewError(twirp.Unknown, err.Error())
	}

	return twirp.NewError(code, err.Error())
}
