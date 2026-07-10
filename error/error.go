package error

import "errors"


var (
	ErrInvalidRequest = errors.New("invalid signup request")
	ErrInvalidEmail   = errors.New("enter valid email")
	ErrUserExists     = errors.New("user already exists")
)