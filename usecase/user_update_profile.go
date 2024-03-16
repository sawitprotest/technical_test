package usecase

import (
	"context"
	"net/http"

	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (u *userUsecaseCtx) UpdateProfile(ctx context.Context, form *user.UpdateProfileRequest, userID int) error {
	var err error
	if err = form.Validation(); err != nil {
		return err
	}

	existsUser, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Internal server error",
		}
	}
	if existsUser == nil {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusForbidden,
			ErrorMessage: "This user does not exists",
		}
	}

	if form.PhoneNumber != `` {
		existsData, err := u.repo.GetUserByPhoneNumber(ctx, form.PhoneNumber)
		if err != nil {
			return &shared.ErrorMessage{
				ErrorCode:    http.StatusUnprocessableEntity,
				ErrorMessage: "Internal server error",
			}
		}
		if existsData != nil {
			return &shared.ErrorMessage{
				ErrorCode:    http.StatusConflict,
				ErrorMessage: "Phone number already exists",
			}
		}
	}

	timeNow := shared.UTC7(u.repo.Now())
	existsUser.PhoneNumber = form.PhoneNumber
	existsUser.FullName = form.FullName
	existsUser.UpdatedAt = &timeNow
	err = u.repo.UpdateProfile(ctx, existsUser)
	if err != nil {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Internal server error",
		}
	}

	return nil
}
