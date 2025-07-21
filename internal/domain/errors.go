package domain

import "errors"

var (
	ErrContentEmpty       = errors.New("content cannot be empty")
	ErrContentTooLong     = errors.New("content exceeds 280 characters")
	ErrUserIDEmpty        = errors.New("userID cannot be empty")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrNoMessagesForUser  = errors.New("user doesn't have any post yet")
	ErrNoFollowersForUser = errors.New("user doesn't have any followers yet")
	ErrNilUserProvided    = errors.New("user should be provided")
	ErrUserAlreadyExists  = errors.New("user already created")
	ErrUserNotFound       = errors.New("user not found")
)
