package api

import (
	"context"

	"github.com/inari111/money-transfer-study/domain/transfer"

	"github.com/inari111/money-transfer-study/proto"
)

func NewMoneyTransferCommand(
	transferRepo transfer.Repository,
) pb.MoneyTransferCommand {
	return &moneyTransferCommand{
		transferRepo: transferRepo,
	}
}

type moneyTransferCommand struct {
	transferRepo transfer.Repository
}

// TODO
func (*moneyTransferCommand) Send(context.Context, *pb.MoneyTransferSendRequest) (*pb.MoneyTransferSendResponse, error) {
	// ユーザー取得

	// 送り元のレコードロック

	// 送り先のレコードロック

	// 送金処理

	panic("implement me")
}
