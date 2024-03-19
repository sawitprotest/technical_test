package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/usecase"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Endpoint for user login.
	// (POST /login)
	Login(ctx echo.Context) error
	// Endpoint for get user profile.
	// (GET /profile/{id})
	GetUserProfile(ctx echo.Context, id string) error
	// Endpoint for update user profile.
	// (PUT /profile/{id})
	UpdateUserProfile(ctx echo.Context, id string) error
	// Endpoint for user registration.
	// (POST /registration)
	Registration(ctx echo.Context) error
}

type handler struct {
	userUsecase usecase.UserUsecase
}

func NewHandler(
	userUsecase usecase.UserUsecase,
) ServerInterface {
	return &handler{
		userUsecase: userUsecase,
	}
}
