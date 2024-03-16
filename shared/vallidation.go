package shared

import (
	"errors"
	"regexp"
)

var (
	PhoneNumberFormat = `^(\+62)[0-9]+$`
)

func CheckPasswordComplexity(pw string) error {
	if len(pw) < 6 || len(pw) > 64 {
		return errors.New(`password length must be between 6 to 64 characters`)
	}
	num := `[0-9]{1}`
	A_Z := `[A-Z]{1}`
	special := `[\W]{1}`

	errorMsg := `password must contains at least 1 uppercase, 1 number, and 1 special characters`

	if b, err := regexp.MatchString(num, pw); !b || err != nil {
		return errors.New(errorMsg)
	}
	if b, err := regexp.MatchString(A_Z, pw); !b || err != nil {
		return errors.New(errorMsg)
	}
	if b, err := regexp.MatchString(special, pw); !b || err != nil {
		return errors.New(errorMsg)
	}

	return nil
}
