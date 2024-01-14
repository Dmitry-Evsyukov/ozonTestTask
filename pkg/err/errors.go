package httpError

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	NoOriginalUrl       = errors.New("original url for your short version wasn't found")
	BadRequest          = errors.New("bad Request")
	InternalServerError = errors.New("unknown internal server error")
)

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func NewRestError(status int, err string, causes interface{}) RestError {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func NewInternalServerError(err error) RestErr {
	return NewRestError(http.StatusInternalServerError, err.Error(), err)
}

func NewBadRequestError(err error) RestErr {
	return NewRestError(http.StatusBadRequest, err.Error(), err)
}

func parseValidatorError(err error) RestErr {
	if strings.Contains(err.Error(), "Url") {
		return NewRestError(http.StatusBadRequest, "Invalid url", err)
	}
	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
}

func ParseError(err error) RestErr {
	switch {
	case strings.Contains(err.Error(), "no rows in result set"):
		return NewRestError(http.StatusNotFound, NoOriginalUrl.Error(), err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
	case strings.Contains(err.Error(), "Field validation"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "short version wasn't found"):
		return NewRestError(http.StatusNotFound, NoOriginalUrl.Error(), err)
	default:
		if re, ok := err.(RestErr); ok {
			return re
		}
		return NewInternalServerError(err)
	}
}

func HandleError(err error) (int, interface{}) {
	return ParseError(err).Status(), ParseError(err)
}
