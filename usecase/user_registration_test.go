package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/sawitpro/technical_test/entity"
	"github.com/sawitpro/technical_test/repository/mocks"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseCtx_UserRegistration(t *testing.T) {
	type args struct {
		form *user.UserRegistrationRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *user.UserRegistrationResponse
		wantErr bool
		err     error
		before  func() *userUsecaseCtx
	}{
		{
			name: "TestUserRegistration-PhoneNumberEmpty",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: ``,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Phone number length must be between 10 to 13 characters",
			},
			before: func() *userUsecaseCtx {

				uc := &userUsecaseCtx{}

				return uc
			},
		},
		{
			name: "TestUserRegistration-PhoneNumberInvalid",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `085812345678`,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Invalid phone number format",
			},
			before: func() *userUsecaseCtx {

				uc := &userUsecaseCtx{}

				return uc
			},
		},
		{
			name: "TestUserRegistration-FullNameEmpty",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `+621234567890`,
					FullName:    ``,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Full name length must be between 3 to 60 characters",
			},
			before: func() *userUsecaseCtx {

				uc := &userUsecaseCtx{}

				return uc
			},
		},
		{
			name: "TestUserRegistration-PasswordInvalid",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    `password`,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "password must contains at least 1 uppercase, 1 number, and 1 special characters",
			},
			before: func() *userUsecaseCtx {

				uc := &userUsecaseCtx{}

				return uc
			},
		},
		{
			name: "TestUserRegistration-GetUserByPhoneError",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    `Password123!`,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusUnprocessableEntity,
				ErrorMessage: "Internal server error",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				uc := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+621234567890`).Return(nil, errors.New("error")).Once()

				return uc
			},
		},
		{
			name: "TestUserRegistration-PhoneNumberExists",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    `Password123!`,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusConflict,
				ErrorMessage: "Phone number already exists",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				uc := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+621234567890`).Return(&entity.User{}, nil).Once()

				return uc
			},
		},
		{
			name: "TestUserRegistration-CreateUserFailed",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    `Password123!`,
				},
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusUnprocessableEntity,
				ErrorMessage: "Internal server error",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				uc := &userUsecaseCtx{
					repo: mockRepo,
				}

				timeNow := time.Now()
				now := shared.UTC7(timeNow)

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+621234567890`).Return(nil, nil).Once()

				mockRepo.On(`Now`).Return(timeNow).Once()

				mockRepo.On(`RandomString`, 12).Return(`123456789ABC`).Once()

				password := shared.MD5(`Password123!` + `123456789ABC`)
				mockUserData := &entity.User{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    password,
					AccountSalt: `123456789ABC`,
					CreatedAt:   now,
					UpdatedAt:   &now,
				}
				mockRepo.On(`Create`, mock.Anything, mockUserData).Return(errors.New(`error`)).Once()

				return uc
			},
		},
		{
			name: "TestUserRegistration-Success",
			args: args{
				form: &user.UserRegistrationRequest{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    `Password123!`,
				},
			},
			want: &user.UserRegistrationResponse{
				UserID: 0,
			},
			wantErr: false,
			err:     nil,
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				uc := &userUsecaseCtx{
					repo: mockRepo,
				}

				timeNow := time.Now()
				now := shared.UTC7(timeNow)

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+621234567890`).Return(nil, nil).Once()

				mockRepo.On(`Now`).Return(timeNow).Once()

				mockRepo.On(`RandomString`, 12).Return(`123456789ABC`).Once()

				password := shared.MD5(`Password123!` + `123456789ABC`)
				mockUserData := &entity.User{
					PhoneNumber: `+621234567890`,
					FullName:    `User123`,
					Password:    password,
					AccountSalt: `123456789ABC`,
					CreatedAt:   now,
					UpdatedAt:   &now,
				}
				mockRepo.On(`Create`, mock.Anything, mockUserData).Return(nil).Once()

				return uc
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.before()
			got, err := u.UserRegistration(context.Background(), tt.args.form)
			if (err != nil) != tt.wantErr || !assert.Equal(t, err, tt.err) {
				t.Errorf("userUsecaseCtx.UserRegistration() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecaseCtx.UserRegistration() = %v, want %v", got, tt.want)
			}
		})
	}
}
