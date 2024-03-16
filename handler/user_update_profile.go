package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/shared"
	"github.com/sawitpro/technical_test/usecase/user"
)

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return shared.HttpError(c, err)
	}

	form := new(user.UpdateProfileRequest)
	if err := c.Bind(form); err != nil {
		return shared.HttpError(c, err)
	}

	err = h.userUsecase.UpdateProfile(ctx, form, userID)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, nil)
	return c.JSON(http.StatusOK, res)
}
