package rpc

import (
	"context"

	"github.com/inari111/money-transfer/proto"
)

func NewUserQuery() pb.UserQuery {
	return &userQuery{}
}

type userQuery struct{}

// TODO
func (*userQuery) Get(context.Context, *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	return &pb.UserGetResponse{
		User: &pb.User{
			Id:          "11111111111111111111111",
			UserProfile: nil,
			CreatedAt:   nil,
			UpdatedAt:   nil,
		},
	}, nil
}
