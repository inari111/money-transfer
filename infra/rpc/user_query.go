package rpc

import (
	"context"

	"github.com/inari111/money-transfer/rpc"

	"github.com/inari111/money-transfer/domain/user"

	userApp "github.com/inari111/money-transfer/application/user"

	"github.com/inari111/money-transfer/proto"
)

func NewUserQuery(
	userApp userApp.Application,
) pb.UserQuery {
	return &userQuery{
		userApp: userApp,
	}
}

type userQuery struct {
	userApp userApp.Application
}

func (q *userQuery) Get(ctx context.Context, req *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	u, err := q.userApp.Get(ctx, user.ID(req.UserId))
	if err != nil {
		return nil, rpc.TwirpErrFrom(err)
	}
	return &pb.UserGetResponse{
		User: &pb.User{
			Id:          uint64(u.ID.Int64()),
			UserProfile: userProfileFrom(u.Profile),
			CreatedAt:   TimestampFrom(u.CreatedAt),
			UpdatedAt:   TimestampFrom(u.UpdatedAt),
		},
	}, nil
}

func userProfileFrom(p *user.Profile) *pb.Profile {
	return &pb.Profile{
		Name: p.Name,
		Age:  uint32(p.Age),
	}
}
