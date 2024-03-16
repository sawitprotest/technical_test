package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (h *UserHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	form := new(user.UserLoginRequest)
	if err := c.Bind(form); err != nil {
		return shared.HttpError(c, err)
	}

	result, err := h.userUsecase.UserLogin(ctx, form)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, result)
	return c.JSON(http.StatusOK, res)
}
