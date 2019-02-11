package rpc

import (
	"context"

	"github.com/twitchtv/twirp"
)

func UserAuthHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		// TODO
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		},
	}
}

func LoggingErrorHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		// TODO
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			return ctx
		},
	}
}
