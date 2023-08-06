package apperror

import (
	"fmt"
	"net/http"
)

type ErrorCode string

type AppError struct {
	StatusCode     int       `json:"-"`
	ErrorCode      ErrorCode `json:"errorCode"`
	Message        any       `json:"message"`
	Detail         any       `json:"detail"`
	internalDetail error     `json:"-"`
	internalDigest string    `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("status code: %d, error code: %s,  external message: %s, external detail: %v, ierr error digest: %s, internal detail: %v", e.StatusCode, e.ErrorCode, e.Message, e.Detail, e.internalDigest, e.internalDetail)
}

// func NewAppError(sc int, ec string, message string, detail interface{}) *AppError {
// 	return &AppError{
// 		StatusCode: sc,
// 		ErrorCode:  ec,
// 		Message:    message,
// 		Detail:     detail,
// 	}
// }

const (
	// 1000-1999: unknown system error
	UnknownError  ErrorCode = "1000"
	EchoHttpError ErrorCode = "1001"
	// 2000-2999: user error
	InvalidRequest       ErrorCode = "2000"
	InvalidRequestBody   ErrorCode = "2001"
	InvalidRequestPath   ErrorCode = "2002"
	InvalidRequestQuery  ErrorCode = "2003"
	InvalidRequestHeader ErrorCode = "2004"
	InvalidRequestCookie ErrorCode = "2005"
	InvalidRequestForm   ErrorCode = "2006"

	NotFound ErrorCode = "2900"

	//  3000-3999: validation error

	// 4000-4999: auth error
	Unauthorized ErrorCode = "4000"
	TokenExpired ErrorCode = "4001"

	// 5000-5999: database error

	// 6000-6999: third party error

	// 7000-7999: ierr error
	InternalError ErrorCode = "7000"
)

func NewUnknownError(ierr error, detail any) *AppError {
	return &AppError{
		StatusCode:     http.StatusInternalServerError,
		ErrorCode:      UnknownError,
		Message:        http.StatusText(http.StatusInternalServerError),
		Detail:         detail,
		internalDetail: ierr,
	}
}
func NewEchoHttpError(code int, message any, ierr error) *AppError {
	return &AppError{
		StatusCode:     code,
		ErrorCode:      EchoHttpError,
		Message:        message,
		internalDetail: ierr,
	}
}

func NewInvalidRequestBodyError(ierr error, detail any) *AppError {
	return &AppError{
		StatusCode:     http.StatusBadRequest,
		ErrorCode:      InvalidRequestBody,
		Message:        "invalid request body",
		Detail:         detail,
		internalDetail: ierr,
	}
}

func NewInvalidRequestPathError(ierr error, detail any) *AppError {
	return &AppError{
		StatusCode:     http.StatusBadRequest,
		ErrorCode:      InvalidRequestPath,
		Message:        "invalid request path",
		Detail:         detail,
		internalDetail: ierr,
	}
}

func NewNotFoundError(ierr error, detail any) *AppError {
	return &AppError{
		StatusCode:     http.StatusNotFound,
		ErrorCode:      NotFound,
		Message:        http.StatusText(http.StatusNotFound),
		Detail:         detail,
		internalDetail: ierr,
	}
}

func NewUnauthorizedError(ierr error, detail any, idg string) *AppError {
	return &AppError{
		StatusCode:     http.StatusUnauthorized,
		ErrorCode:      Unauthorized,
		Message:        http.StatusText(http.StatusUnauthorized),
		Detail:         detail,
		internalDetail: ierr,
		internalDigest: idg,
	}
}

func NewTokenExpiredError(ierr error, detail any) *AppError {
	return &AppError{
		StatusCode:     http.StatusUnauthorized,
		ErrorCode:      TokenExpired,
		Message:        http.StatusText(http.StatusUnauthorized),
		Detail:         detail,
		internalDetail: ierr,
	}
}
func NewInternalError(ierr error, detail any, idg string) *AppError {
	return &AppError{
		StatusCode:     http.StatusInternalServerError,
		ErrorCode:      InternalError,
		Message:        http.StatusText(http.StatusInternalServerError),
		Detail:         detail,
		internalDetail: ierr,
		internalDigest: idg,
	}
}
