package user

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/sawitpro/technical_test/shared"
)

type UserRegistrationRequest struct {
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
}

type UserRegistrationResponse struct {
	UserID int `json:"user_id"`
}

func (c *UserRegistrationRequest) Validation() error {

	if len(c.PhoneNumber) < 10 || len(c.PhoneNumber) > 13 {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Phone number length must be between 10 to 13 characters",
		}
	}

	match, err := regexp.MatchString(shared.PhoneNumberFormat, c.PhoneNumber)
	if err != nil || !match {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Invalid phone number format",
		}
	}

	if len(c.FullName) < 3 || len(c.FullName) > 60 {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Full name length must be between 3 to 60 characters",
		}
	}

	c.Password = strings.TrimSpace(c.Password)
	err = shared.CheckPasswordComplexity(c.Password)
	if err != nil {
		return &shared.ErrorMessage{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
	}

	return nil
}
