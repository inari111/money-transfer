package api

import (
	"context"

	"github.com/inari111/money-transfer/proto"
)

func NewMoneyCommand() pb.MoneyCommand {
	return &moneyCommand{}
}

type moneyCommand struct{}

func (*moneyCommand) Gain(context.Context, *pb.MoneyGainRequest) (*pb.MoneyGainResponse, error) {
	panic("implement me")
}
