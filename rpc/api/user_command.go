package api

import (
	"context"

	pb "github.com/inari111/money-transfer-study/proto"

	"github.com/inari111/money-transfer-study/application/user"
)

func NewUserCommand(
	userApp user.Application,
) pb.UserCommand {
	return &userCommand{
		userApp: userApp,
	}
}

type userCommand struct {
	userApp user.Application
}

func (c *userCommand) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	if err := c.userApp.Register(ctx); err != nil {
		// TODO: twirp errorへ変換
		return nil, err
	}
	return &pb.UserRegisterResponse{}, nil
}

func (c *userCommand) UpdateProfile(ctx context.Context, req *pb.UserUpdateProfileRequest) (*pb.UserUpdateProfileResponse, error) {
	return &pb.UserUpdateProfileResponse{}, nil
}
