package usecase

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sawitpro/technical_test/config"
	"github.com/sawitpro/technical_test/entity"
	"github.com/sawitpro/technical_test/repository/mocks"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseCtx_UserLogin(t *testing.T) {
	type args struct {
		form *user.UserLoginRequest
	}

	mockUserData := &entity.User{
		ID:          1,
		Password:    `d1a7c9c2fba028ce3899850143ab504f`,
		AccountSalt: `SALT_STRING`,
	}
	timeNow := time.Now()
	privateKey := mockInitPrivateKey()
	loginResponse := mockCreateAccessToken(mockUserData, privateKey, timeNow)

	tests := []struct {
		name    string
		args    args
		want    *user.UserLoginResponse
		wantErr bool
		err     error
		before  func() *userUsecaseCtx
	}{
		{
			name: `TestLogin-PhoneNumberEmpty`,
			args: args{
				form: &user.UserLoginRequest{
					PhoneNumber: ``,
				},
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Phone number is required",
			},
			before: func() *userUsecaseCtx {
				u := &userUsecaseCtx{}

				return u
			},
		},
		{
			name: `TestLogin-PasswordEmpty`,
			args: args{
				form: &user.UserLoginRequest{
					PhoneNumber: `+62123456789`,
				},
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Password is required",
			},
			before: func() *userUsecaseCtx {
				u := &userUsecaseCtx{}

				return u
			},
		},
		{
			name: `TestLogin-PhoneNumberNotExists`,
			args: args{
				form: &user.UserLoginRequest{
					PhoneNumber: `+62123456789`,
					Password:    `Password123!`,
				},
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "This phone number is not registered",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(nil, nil).Once()

				return u
			},
		},
		{
			name: `TestLogin-WrongPassword`,
			args: args{
				form: &user.UserLoginRequest{
					PhoneNumber: `+62123456789`,
					Password:    `Password123!`,
				},
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Wrong password",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockUserData := &entity.User{
					Password:    `20404d0af589a9ac11f8043dc0128674`,
					AccountSalt: `SALT_STRING`,
				}
				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(mockUserData, nil).Once()

				return u
			},
		},
		{
			name: `TestLogin-IncrementLoginError`,
			args: args{
				form: &user.UserLoginRequest{
					PhoneNumber: `+62123456789`,
					Password:    `Password123!`,
				},
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusUnprocessableEntity,
				ErrorMessage: "Internal server error",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)
				cfg := &config.Config{}
				cfg.PrivateKey = privateKey

				u := &userUsecaseCtx{
					repo: mockRepo,
					cfg:  cfg,
				}

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(mockUserData, nil).Once()

				mockRepo.On(`Now`).Return(timeNow)

				mockRepo.On(`IncrementSuccessfulLogin`, mock.Anything, 1).Return(errors.New(`error`)).Once()

				return u
			},
		},
		{
			name: `TestLogin-Success`,
			args: args{
				form: &user.UserLoginRequest{
					PhoneNumber: `+62123456789`,
					Password:    `Password123!`,
				},
			},
			want:    loginResponse,
			wantErr: false,
			err:     nil,
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)
				cfg := &config.Config{}
				cfg.PrivateKey = privateKey

				u := &userUsecaseCtx{
					repo: mockRepo,
					cfg:  cfg,
				}

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(mockUserData, nil).Once()

				mockRepo.On(`Now`).Return(timeNow)

				mockRepo.On(`IncrementSuccessfulLogin`, mock.Anything, 1).Return(nil).Once()

				return u
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.before()
			got, err := u.UserLogin(context.Background(), tt.args.form)
			if (err != nil) != tt.wantErr || !assert.Equal(t, err, tt.err) {
				t.Errorf("userUsecaseCtx.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecaseCtx.UserLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockCreateAccessToken(data *entity.User, privateKey *rsa.PrivateKey, now time.Time) *user.UserLoginResponse {
	claim := entity.AccessTokenClaim{
		UserID:      data.ID,
		PhoneNumber: data.PhoneNumber,
	}

	end := now.Add(time.Hour)

	claim.IssuedAt = jwt.NewNumericDate(now)
	claim.ExpiresAt = jwt.NewNumericDate(end)

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, _ := newToken.SignedString(privateKey)

	res := &user.UserLoginResponse{
		UserID:    data.ID,
		Token:     tokenString,
		ExpiredAt: end.Format(time.RFC3339),
	}
	return res
}

func mockInitPrivateKey() *rsa.PrivateKey {
	keyPath := `../config/app.rsa`
	signBytes, err := os.ReadFile(keyPath)
	fmt.Println(err)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fmt.Println(err)
	return signKey
}
