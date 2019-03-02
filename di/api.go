// +build wireinject

package di

import (
	"net/http"

	userApp "github.com/inari111/money-transfer/application/user"

	"github.com/inari111/money-transfer/domain"

	"github.com/inari111/money-transfer/infra/persistence/mysql"

	"github.com/inari111/money-transfer/handler"

	"github.com/inari111/money-transfer/infra/persistence/repository"
	infra_rpc "github.com/inari111/money-transfer/infra/rpc"

	"github.com/google/wire"
	"github.com/inari111/money-transfer/rpc/api"
)

func InitializeAPIHandler() http.Handler {
	wire.Build(
		mysql.NewDbMap,

		repository.NewUserRepository,
		repository.NewUserProfileRepository,
		repository.NewTransferRepository,

		domain.NewCurrentTimeFunc,

		userApp.NewApplication,

		api.NewMoneyTransferCommand,
		api.NewUserCommand,
		api.NewMoneyCommand,

		infra_rpc.NewMoneyQuery,
		infra_rpc.NewUserQuery,

		handler.NewHandler,
	)
	return nil
}
