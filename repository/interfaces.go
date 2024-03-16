package repository

import (
	"context"
	"time"

	"github.com/sawitpro/technical_test/entity"
)

type Repository interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error)
	GetUserByID(ctx context.Context, userID int) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	IncrementSuccessfulLogin(ctx context.Context, userID int) error
	UpdateProfile(ctx context.Context, user *entity.User) error

	Now() time.Time
	RandomString(length int) string
}
