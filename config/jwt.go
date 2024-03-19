package config

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/entity"
)

func JWTVerify(rsaPublicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			header := req.Header
			auth := header.Get("Authorization")

			if len(auth) <= 0 {
				return echo.NewHTTPError(http.StatusForbidden, "authorization is empty")
			}

			splitToken := strings.Split(auth, " ")
			if len(splitToken) < 2 {
				return echo.NewHTTPError(http.StatusForbidden, "authorization is empty")
			}

			if splitToken[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusForbidden, "authorization is invalid")
			}

			tokenStr := splitToken[1]
			token, err := jwt.ParseWithClaims(tokenStr, &entity.AccessTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return rsaPublicKey, nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			if claims, ok := token.Claims.(*entity.AccessTokenClaim); token.Valid && ok {
				c.Set("UserID", claims.UserID)
				c.Set("PhoneNumber", claims.PhoneNumber)

				return next(c)
			}

			return echo.NewHTTPError(http.StatusForbidden, "Unknown token error")
		}
	}
}
