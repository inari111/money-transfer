package handler

import (
	"net/http"

	pb "github.com/inari111/money-transfer/proto"
	"github.com/inari111/money-transfer/rpc"

	"github.com/twitchtv/twirp"
)

type server struct {
	pathPrefix string
	twServer   pb.TwirpServer
}

func NewHandler(
	moneyTransferCommand pb.MoneyTransferCommand,
	moneyQuery pb.MoneyQuery,
	userCommand pb.UserCommand,
	userQuery pb.UserQuery,
	moneyCommand pb.MoneyCommand,
) http.Handler {
	mux := http.NewServeMux()

	hooks := twirp.ChainHooks(
		rpc.UserAuthHooks(),
		rpc.LoggingErrorHooks(),
	)

	var servers = []server{
		{
			pathPrefix: pb.MoneyTransferCommandPathPrefix,
			twServer: pb.NewMoneyTransferCommandServer(
				moneyTransferCommand,
				hooks,
			),
		},
		{
			pathPrefix: pb.MoneyQueryPathPrefix,
			twServer: pb.NewMoneyQueryServer(
				moneyQuery,
				hooks,
			),
		},
		{
			pathPrefix: pb.UserCommandPathPrefix,
			twServer: pb.NewUserCommandServer(
				userCommand,
				hooks,
			),
		},
		{
			pathPrefix: pb.UserQueryPathPrefix,
			twServer: pb.NewUserQueryServer(
				userQuery,
				hooks,
			),
		},
		{
			pathPrefix: pb.MoneyCommandPathPrefix,
			twServer: pb.NewMoneyCommandServer(
				moneyCommand,
				hooks,
			),
		},
	}
	for _, s := range servers {
		mux.Handle(s.pathPrefix, s.twServer)
	}
	return mux
}
