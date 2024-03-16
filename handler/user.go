package handler

import (
	"github.com/sawitpro/technical_test/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}
