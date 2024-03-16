package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sawitpro/technical_test/entity"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (u *userUsecaseCtx) UserLogin(ctx context.Context, form *user.UserLoginRequest) (*user.UserLoginResponse, error) {
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
	if existsUser == nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "This phone number is not registered",
		}
	}

	hashedPassword := shared.MD5(form.Password + existsUser.AccountSalt)
	if hashedPassword != existsUser.Password {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Wrong password",
		}
	}

	res, err := u.createAccessToken(existsUser)
	if err != nil {
		return nil, err
	}

	err = u.repo.IncrementSuccessfulLogin(ctx, existsUser.ID)
	if err != nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Internal server error",
		}
	}

	return res, nil
}

func (u *userUsecaseCtx) createAccessToken(data *entity.User) (*user.UserLoginResponse, error) {
	var err error

	claim := entity.AccessTokenClaim{
		UserID:      data.ID,
		PhoneNumber: data.PhoneNumber,
	}

	now := u.repo.Now()
	end := now.Add(time.Hour)

	claim.IssuedAt = jwt.NewNumericDate(now)
	claim.ExpiresAt = jwt.NewNumericDate(end)

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, err := newToken.SignedString(u.cfg.PrivateKey)
	if err != nil {
		return nil, &shared.ErrorMessage{
			ErrorCode:    http.StatusUnprocessableEntity,
			ErrorMessage: "Failed sign access token",
		}
	}

	res := &user.UserLoginResponse{
		UserID:    data.ID,
		Token:     tokenString,
		ExpiredAt: end.Format(time.RFC3339),
	}
	return res, nil
}
