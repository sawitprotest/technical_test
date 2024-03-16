package usecase

import (
	"context"
	"net/http"

	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (u *userUsecaseCtx) GetUserProfile(ctx context.Context, userID int) (*user.GetUserProfileResponse, error) {
	var (
		err error
		res *user.GetUserProfileResponse
	)

	existsUser, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Internal server error",
		}
	}
	if existsUser == nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusForbidden,
			ErrorMessage: "This user does not exists",
		}
	}

	res = &user.GetUserProfileResponse{
		FullName:    existsUser.FullName,
		PhoneNumber: existsUser.PhoneNumber,
	}

	return res, nil
}
