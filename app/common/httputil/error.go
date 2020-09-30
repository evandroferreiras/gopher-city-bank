package httputil

import "net/http"

// HTTPErrorParseBody is the http error returned when got a parse body error
var HTTPErrorParseBody = HTTPError{
	Code:    http.StatusBadRequest,
	Message: "error when trying to parse body",
}

// HTTPErrorValidateBody is the http error returned when got a validate body error. For instance, required field is missing
var HTTPErrorValidateBody = HTTPError{
	Code:    http.StatusBadRequest,
	Message: "error when trying to validate body. Check the required fields",
}

// NewError is a generic http error
func NewError(status int, err error) HTTPError {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return er
}

// NewHTTPErrorValidateBody is a constructor to validatebody errors
func NewHTTPErrorValidateBody(err error) HTTPError {
	httpErr := HTTPErrorValidateBody
	httpErr.Error = err.Error()

	return httpErr
}

// HTTPError is a constructor to generic http errors.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
	Error   string `json:"error,omitempty"  example:"status bad request"`
}
