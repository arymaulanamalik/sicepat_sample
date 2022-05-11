package shared

import (
	"strconv"
	"strings"

	"gitlab.sicepat.tech/platform/golib/response"
)

type ErrorResponse struct {
	Type    error
	Message string
	Key     string
}

func GetErrorResponse(err error) (res ErrorResponse) {
	errSplit := strings.Split(err.Error(), " ~~ ")

	res.Type = response.ErrInternalServerError
	res.Key = "[INTERNAL]"
	if len(errSplit) == 1 {
		res.Message = err.Error()
		return
	}

	errCode, _ := strconv.Atoi(errSplit[0])
	switch errCode {
	case 400:
		res.Type = response.ErrBadRequest
		res.Key = "[REQUEST]"
	case 401:
		res.Type = response.ErrUnauthorized
		res.Key = "[TOKEN]"
	}

	res.Message = errSplit[1]

	return
}
