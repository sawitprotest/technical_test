package usecase

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/sawitpro/technical_test/entity"
	"github.com/sawitpro/technical_test/repository/mocks"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseCtx_GetUserProfile(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    *user.GetUserProfileResponse
		wantErr bool
		err     error
		before  func() *userUsecaseCtx
	}{
		{
			name: `TestGetProfile-GetUserByIDError`,
			args: args{
				userID: 1,
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

				mockRepo.On(`GetUserByID`, mock.Anything, 1).Return(nil, errors.New("error")).Once()

				return uc
			},
		},
		{
			name: `TestGetProfile-UserNotExists`,
			args: args{
				userID: 1,
			},
			want:    nil,
			wantErr: true,
			err: &shared.ErrorMessage{
				ErrorCode:    http.StatusForbidden,
				ErrorMessage: "This user does not exists",
			},
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				uc := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockRepo.On(`GetUserByID`, mock.Anything, 1).Return(nil, nil).Once()

				return uc
			},
		},
		{
			name: `TestGetProfile-Success`,
			args: args{
				userID: 1,
			},
			want: &user.GetUserProfileResponse{
				FullName:    `User123`,
				PhoneNumber: `+62123456789`,
			},
			wantErr: false,
			err:     nil,
			before: func() *userUsecaseCtx {
				mockRepo := new(mocks.Repository)

				uc := &userUsecaseCtx{
					repo: mockRepo,
				}

				mockUserData := &entity.User{
					ID:          1,
					FullName:    "User123",
					PhoneNumber: `+62123456789`,
				}
				mockRepo.On(`GetUserByID`, mock.Anything, 1).Return(mockUserData, nil).Once()

				return uc
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.before()
			got, err := u.GetUserProfile(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr || !assert.Equal(t, err, tt.err) {
				t.Errorf("userUsecaseCtx.GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecaseCtx.GetUserProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
