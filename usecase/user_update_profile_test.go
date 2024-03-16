package usecase

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/sawitpro/technical_test/entity"
	"github.com/sawitpro/technical_test/repository/mocks"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseCtx_UpdateProfile(t *testing.T) {
	type args struct {
		form   *user.UpdateProfileRequest
		userID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		before  func() *userUsecaseCtx
	}{
		{
			name: "TestUpdateProfile-PhoneNumberInvalid",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `0123456789`,
				},
				userID: 1,
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Invalid phone number format",
			},
			before: func() *userUsecaseCtx {
				u := &userUsecaseCtx{}

				return u
			},
		},
		{
			name: "TestUpdateProfile-FullNameInvalid",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `+62123456789`,
					FullName:    `hi`,
				},
				userID: 1,
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "Full name length must be between 3 to 60 characters",
			},
			before: func() *userUsecaseCtx {
				u := &userUsecaseCtx{}

				return u
			},
		},
		{
			name: "TestUpdateProfile-GetUserError",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
				},
				userID: 1,
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusUnprocessableEntity,
				ErrorMessage: "Internal server error",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On("GetUserByID", mock.Anything, 1).Return(nil, errors.New(`error`)).Once()

				return u
			},
		},
		{
			name: "TestUpdateProfile-GetUserError",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
				},
				userID: 1,
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusForbidden,
				ErrorMessage: "This user does not exists",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On("GetUserByID", mock.Anything, 1).Return(nil, nil).Once()

				return u
			},
		},
		{
			name: "TestUpdateProfile-PhoneNumberExists",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
				},
				userID: 1,
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusConflict,
				ErrorMessage: "Phone number already exists",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On("GetUserByID", mock.Anything, 1).Return(&entity.User{}, nil).Once()

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(&entity.User{}, nil).Once()

				return u
			},
		},
		{
			name: "TestUpdateProfile-UpdateError",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
				},
				userID: 1,
			},
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusConflict,
				ErrorMessage: "Phone number already exists",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				timeNow := time.Now()
				now := shared.UTC7(timeNow)

				mockRepo.On("GetUserByID", mock.Anything, 1).Return(&entity.User{}, nil).Once()

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(nil, nil).Once()

				mockRepo.On(`Now`).Return(timeNow)

				mockUserData := &entity.User{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
					UpdatedAt:   &now,
				}
				mockRepo.On(`UpdateProfile`, mock.Anything, mockUserData).Return(errors.New(`error`)).Once()

				return u
			},
		},
		{
			name: "TestUpdateProfile-UpdateSucccess",
			args: args{
				form: &user.UpdateProfileRequest{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
				},
				userID: 1,
			},
			wantErr: false,
			err:     nil,
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				u := &userUsecaseCtx{
					repo: mockRepo,
				}

				timeNow := time.Now()
				now := shared.UTC7(timeNow)

				mockRepo.On("GetUserByID", mock.Anything, 1).Return(&entity.User{}, nil).Once()

				mockRepo.On(`GetUserByPhoneNumber`, mock.Anything, `+62123456789`).Return(nil, nil).Once()

				mockRepo.On(`Now`).Return(timeNow)

				mockUserData := &entity.User{
					PhoneNumber: `+62123456789`,
					FullName:    `user123`,
					UpdatedAt:   &now,
				}
				mockRepo.On(`UpdateProfile`, mock.Anything, mockUserData).Return(nil).Once()

				return u
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.before()
			if err := u.UpdateProfile(context.Background(), tt.args.form, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("userUsecaseCtx.UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
