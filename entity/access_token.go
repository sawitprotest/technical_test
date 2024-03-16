package entity

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaim struct {
	jwt.RegisteredClaims
	UserID      int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
}
