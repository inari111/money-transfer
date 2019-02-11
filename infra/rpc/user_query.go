package rpc

import (
	"context"

	"github.com/inari111/money-transfer-study/proto"
)

func NewUserQuery() pb.UserQuery {
	return &userQuery{}
}

type userQuery struct{}

// TODO
func (*userQuery) Get(context.Context, *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	return &pb.UserGetResponse{
		User: &pb.User{
			Id:          "",
			UserProfile: nil,
			CreatedAt:   nil,
			UpdatedAt:   nil,
		},
	}, nil
}
