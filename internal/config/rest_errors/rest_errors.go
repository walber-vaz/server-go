package rest_errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Errors  []Err  `json:"errors,omitempty"`
}

type Err struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *RestErr) Error() string {
	return e.Message
}

func NewRestError(message string, status int, errors []Err) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Errors:  errors,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewBadRequestErrorValidation(message string, errors []Err) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Errors:  errors,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusForbidden,
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusConflict,
	}
}
