package api

import (
	"context"

	"github.com/inari111/money-transfer/rpc"

	"github.com/inari111/money-transfer/domain/user"

	pb "github.com/inari111/money-transfer/proto"

	userApp "github.com/inari111/money-transfer/application/user"
)

func NewUserCommand(
	userApp userApp.Application,
) pb.UserCommand {
	return &userCommand{
		userApp: userApp,
	}
}

type userCommand struct {
	userApp userApp.Application
}

func (c *userCommand) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	if err := c.userApp.Register(ctx); err != nil {
		return nil, rpc.TwirpErrFrom(err)
	}
	return &pb.UserRegisterResponse{}, nil
}

func (c *userCommand) UpdateProfile(ctx context.Context, req *pb.UserUpdateProfileRequest) (*pb.UserUpdateProfileResponse, error) {
	// TODO: 認証部分を作成したら、idTokenをdecodeしてuser.IDを取得し、使うようにする
	err := c.userApp.UpdateProfile(
		ctx,
		&user.Profile{
			ID:   "1",
			Name: req.GetName(),
			Age:  int(req.GetAge()),
		})
	if err != nil {
		return nil, rpc.TwirpErrFrom(err)
	}
	return &pb.UserUpdateProfileResponse{}, nil
}
