// +build wireinject

package di

import (
	"net/http"

	"github.com/inari111/money-transfer-study/domain"

	"github.com/inari111/money-transfer-study/infra/persistence/mysql"

	"github.com/inari111/money-transfer-study/handler"

	"github.com/inari111/money-transfer-study/infra/persistence/repository"
	infra_rpc "github.com/inari111/money-transfer-study/infra/rpc"

	"github.com/google/wire"
	"github.com/inari111/money-transfer-study/rpc/api"
)

func InitializeAPIHandler() http.Handler {
	wire.Build(
		mysql.NewDbMap,

		repository.NewUserRepository,
		repository.NewTransferRepository,

		domain.NewCurrentTimeFunc,

		api.NewMoneyTransferCommand,
		api.NewUserCommand,
		api.NewMoneyCommand,

		infra_rpc.NewMoneyQuery,
		infra_rpc.NewUserQuery,

		handler.NewHandler,
	)
	return nil
}