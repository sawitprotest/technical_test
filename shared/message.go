package shared

import (
	"encoding/json"
	"net/http"
)

var (
	ERR_BAD_REQUEST = ErrorMessage{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: "Malformat Request",
	}
)

type ErrorMessage struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"message"`
}

func (c *ErrorMessage) Error() string {
	b, _ := json.Marshal(c)

	return string(b)
}
