package usecase

import (
	"context"

	"github.com/sawitpro/technical_test/config"
	"github.com/sawitpro/technical_test/repository"
	"github.com/sawitpro/technical_test/usecase/user"
)

type UserUsecase interface {
	UserRegistration(ctx context.Context, form *user.UserRegistrationRequest) (*user.UserRegistrationResponse, error)
	UserLogin(ctx context.Context, form *user.UserLoginRequest) (*user.UserLoginResponse, error)
	GetUserProfile(ctx context.Context, userID int) (*user.GetUserProfileResponse, error)
	UpdateProfile(ctx context.Context, form *user.UpdateProfileRequest, userID int) error
}

type userUsecaseCtx struct {
	cfg  *config.Config
	repo repository.Repository
}

func NewUserUsecase(
	cfg *config.Config,
	repo repository.Repository,
) UserUsecase {
	return &userUsecaseCtx{
		cfg:  cfg,
		repo: repo,
	}
}
