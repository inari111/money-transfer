package user

import (
	"context"

	"github.com/inari111/money-transfer-study/domain"

	"github.com/inari111/money-transfer-study/domain/user"
)

type Application interface {
	// TODO: IDはautoincrementにしているため指定しないが、Firebase Authで匿名ユーザーを作成した場合、
	//  Firebase Authのidを使うべき(Registerの引数にidを受け取るようにする)
	Register(ctx context.Context) error

	UpdateProfile(
		ctx context.Context,
		profile *user.Profile,
	) error
}

func NewApplication(
	useRepo user.Repository,
	profileRepo user.ProfileRepository,
	now domain.CurrentTimeFunc,
) Application {
	return &userApp{
		userRepo:    useRepo,
		profileRepo: profileRepo,
		now:         now,
	}
}

type userApp struct {
	userRepo    user.Repository
	profileRepo user.ProfileRepository
	now         domain.CurrentTimeFunc
}

func (a *userApp) Register(ctx context.Context) error {
	u := &user.User{
		ID:        "",
		CreatedAt: a.now(),
		UpdatedAt: a.now(),
	}
	return a.userRepo.Put(ctx, u)
}

func (a *userApp) UpdateProfile(
	ctx context.Context,
	profile *user.Profile,
) error {
	return a.profileRepo.Put(ctx, profile)
}
