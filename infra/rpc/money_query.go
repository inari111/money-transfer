package rpc

import (
	"context"

	"github.com/inari111/money-transfer-study/proto"
)

func NewMoneyQuery() pb.MoneyQuery {
	return &moneyQuery{}
}

type moneyQuery struct{}

// TODO
func (*moneyQuery) GetBalance(context.Context, *pb.MoneyGetBalanceRequest) (*pb.MoneyGetBalanceResponse, error) {
	panic("implement me")
}
