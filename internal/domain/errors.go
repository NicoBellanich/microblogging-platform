package domain

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrContentEmpty         = errors.New("content cannot be empty")
	ErrContentTooLong       = errors.New("content exceeds 280 characters")
	ErrUserNameEmpty        = errors.New("userName cannot be empty")
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrNoMessagesForUser    = errors.New("user doesn't have any post yet")
	ErrNoFollowersForUser   = errors.New("user doesn't have any followers yet")
	ErrNilUserProvided      = errors.New("user should be provided")
	ErrUserAlreadyExists    = errors.New("user already created")
	ErrUserAlreadyFollowing = errors.New("follow already done")
	ErrUserNotFound         = errors.New("user not found")
	ErrMethodNotAllowed     = errors.New("method not allowed")
	ErrInvalidRequestBody   = errors.New("invalid request body")
)

var baseErrorMap = map[error]int{
	ErrContentEmpty:         http.StatusBadRequest,
	ErrContentTooLong:       http.StatusBadRequest,
	ErrUserNameEmpty:        http.StatusBadRequest,
	ErrInvalidArgument:      http.StatusBadRequest,
	ErrNilUserProvided:      http.StatusBadRequest,
	ErrInvalidRequestBody:   http.StatusBadRequest,
	ErrNoMessagesForUser:    http.StatusNotFound,
	ErrNoFollowersForUser:   http.StatusNotFound,
	ErrUserAlreadyExists:    http.StatusConflict,
	ErrUserAlreadyFollowing: http.StatusConflict,
	ErrUserNotFound:         http.StatusNotFound,
	ErrMethodNotAllowed:     http.StatusMethodNotAllowed,
}

type AppError struct {
	Code     int    // HTTP status
	Message  string // API message
	Op       string // context
	Resource string // missing resource (ej. "username=nico")
	Err      error
}

func NewAppError(op string, baseErr error, resource string) *AppError {
	code, ok := baseErrorMap[baseErr]
	if !ok {
		code = http.StatusInternalServerError
	}
	return &AppError{
		Code:     code,
		Message:  baseErr.Error(),
		Op:       op,
		Resource: resource,
		Err:      baseErr,
	}
}

// Error, implements error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Op, e.Message, e.Resource)
}

// Unwrap key to compare errors with some errors libraries
func (e *AppError) Unwrap() error {
	return e.Err
}
