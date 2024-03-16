package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sawitpro/technical_test/shared"
)

func (h *UserHandler) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return shared.HttpError(c, err)
	}

	result, err := h.userUsecase.GetUserProfile(ctx, userID)
	if err != nil {
		return shared.HttpError(c, err)
	}

	res := shared.JSONSuccess(`Success`, result)
	return c.JSON(http.StatusOK, res)
}
