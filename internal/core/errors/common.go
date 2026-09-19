package core_errors

import "errors"

var (
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrNotFound           = errors.New("request not found")
	ErrConflict           = errors.New("http conflict")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidAuthData    = errors.New("invalid email or password")
	ErrGenerateToken      = errors.New("An error occurred while generating the token")
	ErrConnectDB          = errors.New("connection DB error")
	ErrViolatesForeignKey = errors.New("violates foreign key")
)
