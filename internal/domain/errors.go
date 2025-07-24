package domain

import "errors"

var (
	ErrContentEmpty         = errors.New("content cannot be empty")
	ErrContentTooLong       = errors.New("content exceeds 280 characters")
	ErrUserNameEmpty        = errors.New("userName cannot be empty")
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrNoMessagesForUser    = errors.New("user doesn't have any post yet")
	ErrNoFollowersForUser   = errors.New("user doesn't have any followers yet")
	ErrNilUserProvided      = errors.New("user should be provided")
	ErrUserAlreadyExists    = errors.New("user already created")
	ErrUserAlreadyFollowing = errors.New("user already follow this user")
	ErrUserNotFound         = errors.New("user not found")
)
