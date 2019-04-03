package api

import (
	"context"
	"fmt"

	"github.com/inari111/money-transfer/domain/errors"

	userApp "github.com/inari111/money-transfer/application/user"

	"github.com/inari111/money-transfer/domain/user"

	"github.com/inari111/money-transfer/domain/money"

	"github.com/inari111/money-transfer/rpc"

	moneyApp "github.com/inari111/money-transfer/application/money"

	"github.com/inari111/money-transfer/proto"
)

func NewMoneyCommand(
	moneyApp moneyApp.Application,
	userApp userApp.Application,
) pb.MoneyCommand {
	return &moneyCommand{
		moneyApp: moneyApp,
		userApp:  userApp,
	}
}

type moneyCommand struct {
	moneyApp moneyApp.Application
	userApp  userApp.Application
}

func (c *moneyCommand) Gain(ctx context.Context, req *pb.MoneyGainRequest) (*pb.MoneyGainResponse, error) {
	userID := user.ID(req.GetUserId())
	isExisting, err := c.userApp.Exist(ctx, userID)
	if err != nil {
		return nil, rpc.TwirpErrFrom(err)
	}
	if !isExisting {
		return nil,
			rpc.TwirpErrFrom(
				errors.New(
					fmt.Sprintf("userID:%v is not found", userID),
					errors.NotFound,
				),
			)
	}

	err = c.moneyApp.Gain(
		ctx,
		&money.Money{
			UserID: user.ID(req.GetUserId()),
			Amount: uint64(req.GetAmount()),
		},
	)
	if err != nil {
		return nil, rpc.TwirpErrFrom(err)
	}

	return &pb.MoneyGainResponse{}, nil
}
