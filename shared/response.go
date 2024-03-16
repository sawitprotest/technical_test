package shared

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HttpError(c echo.Context, err error) error {
	switch err.(type) {
	case *ErrorMessage:
		msg := err.(*ErrorMessage)
		return c.JSON(msg.ErrorCode, msg)
	}

	return c.JSON(http.StatusBadRequest, ERR_BAD_REQUEST)
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONSuccess(message string, data interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
