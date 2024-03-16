package usecase

import (
	"context"
	"net/http"

	"github.com/sawitpro/technical_test/entity"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (u *userUsecaseCtx) UserRegistration(ctx context.Context, form *user.UserRegistrationRequest) (*user.UserRegistrationResponse, error) {
	var err error
	if err = form.Validation(); err != nil {
		return nil, err
	}

	existsUser, err := u.repo.GetUserByPhoneNumber(ctx, form.PhoneNumber)
	if err != nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Internal server error",
		}
	}
	if existsUser != nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusConflict,
			ErrorMessage: "Phone number already exists",
		}
	}

	accountSalt := u.repo.RandomString(12)
	hashedPassword := shared.MD5(form.Password + accountSalt)

	timeNow := shared.UTC7(u.repo.Now())
	userData := &entity.User{
		FullName:    form.FullName,
		PhoneNumber: form.PhoneNumber,
		Password:    hashedPassword,
		AccountSalt: accountSalt,
		CreatedAt:   timeNow,
		UpdatedAt:   &timeNow,
	}

	err = u.repo.Create(ctx, userData)
	if err != nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Internal server error",
		}
	}

	res := &user.UserRegistrationResponse{
		UserID: userData.ID,
	}
	return res, nil
}
