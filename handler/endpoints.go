package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (h *handler) Login(c echo.Context) error {
	reqCtx := c.Request().Context()

	form := new(user.UserLoginRequest)
	if err := c.Bind(form); err != nil {
		return shared.HttpError(c, err)
	}

	result, err := h.userUsecase.UserLogin(reqCtx, form)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, result)
	return c.JSON(http.StatusOK, res)
}

func (h *handler) GetUserProfile(c echo.Context, id string) error {
	reqCtx := c.Request().Context()

	userID, err := strconv.Atoi(id)
	if err != nil {
		return shared.HttpError(c, err)
	}

	result, err := h.userUsecase.GetUserProfile(reqCtx, userID)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, result)
	return c.JSON(http.StatusOK, res)
}

func (h *handler) UpdateUserProfile(c echo.Context, id string) error {
	reqCtx := c.Request().Context()

	userID, err := strconv.Atoi(id)
	if err != nil {
		return shared.HttpError(c, err)
	}

	form := new(user.UpdateProfileRequest)
	if err := c.Bind(form); err != nil {
		return shared.HttpError(c, err)
	}

	err = h.userUsecase.UpdateProfile(reqCtx, form, userID)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, nil)
	return c.JSON(http.StatusOK, res)
}

func (h *handler) Registration(c echo.Context) error {
	reqCtx := c.Request().Context()

	form := new(user.UserRegistrationRequest)
	if err := c.Bind(form); err != nil {
		return shared.HttpError(c, err)
	}

	result, err := h.userUsecase.UserRegistration(reqCtx, form)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, result)
	return c.JSON(http.StatusOK, res)
}
