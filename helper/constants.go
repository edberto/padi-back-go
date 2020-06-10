package helper

import "errors"

var (
	ErrUserExisted  = errors.New("User already existed!")
	ErrUserNotFound = errors.New("User not found!")
	ErrTokenExpired = errors.New("Token Expired!")
)
