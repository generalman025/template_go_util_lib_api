package rest_errors

import (
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	err     string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.message
}
func (e restErr) Status() int {
	return e.status
}
func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - cause: [%v]",
		e.message, e.status, e.err, e.causes)
}
func (e restErr) Causes() []interface{} {
	return e.causes
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		status:  status,
		err:     err,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		err:     "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		err:     "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: "unable to retrieve user information from given access_token",
		status:  http.StatusUnauthorized,
		err:     "unauthorized",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status:  http.StatusInternalServerError,
		err:     "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
