package user

import (
	"net/http"

	"github.com/sawitpro/technical_test/shared"
)

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserLoginResponse struct {
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

func (c *UserLoginRequest) Validation() error {

	if c.PhoneNumber == `` {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Phone number is required",
		}
	}

	if c.Password == `` {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Password is required",
		}
	}

	return nil
}
