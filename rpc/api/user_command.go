package api

import (
	"context"

	"github.com/inari111/money-transfer-study/domain"

	"github.com/inari111/money-transfer-study/domain/user"
	"github.com/inari111/money-transfer-study/proto"
)

func NewUserCommand(
	userRepo user.Repository,
	now domain.CurrentTimeFunc,
) pb.UserCommand {
	return &userCommand{
		userRepo: userRepo,
		now:      now,
	}
}

type userCommand struct {
	userRepo user.Repository
	now      domain.CurrentTimeFunc
}

func (c *userCommand) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	u := &user.User{
		ID:        "",
		CreatedAt: c.now(),
		UpdatedAt: c.now(),
	}
	if err := c.userRepo.Put(ctx, u); err != nil {
		// TODO: twirp errへの変換
		return nil, err
	}
	return &pb.UserRegisterResponse{}, nil
}

// TODO
func (*userCommand) UpdateProfile(context.Context, *pb.UserUpdateProfileRequest) (*pb.UserUpdateProfileResponse, error) {
	panic("implement me")
}
