package errors

import (
	"errors"
)

var (
	ErrTimeout       = errors.New("deadline exceeded")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrParseJWTToken = errors.New("error parse jwt token")
	ErrAddSignKey    = errors.New("errors add sign key")
)
