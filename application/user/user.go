package user

import (
	"context"

	"github.com/inari111/money-transfer/domain/errors"

	"github.com/inari111/money-transfer/domain"

	"github.com/inari111/money-transfer/domain/user"
)

type Application interface {
	Get(ctx context.Context, userID user.ID) (*user.User, error)

	Exist(
		ctx context.Context,
		userID user.ID,
	) (bool, error)

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

func (a *userApp) Get(ctx context.Context, userID user.ID) (*user.User, error) {
	return a.userRepo.Get(ctx, userID)
}

func (a *userApp) Exist(
	ctx context.Context,
	userID user.ID,
) (bool, error) {
	_, err := a.userRepo.Get(ctx, userID)
	if err != nil {
		if errors.ErrTypeFrom(err) == errors.NotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (a *userApp) Register(ctx context.Context) error {
	u := &user.User{
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
