package errors

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrDb            = errors.New("database error")
	ErrValidation    = errors.New("validation failed")
	ErrAuth          = errors.New("wrong login or password")
	ErrBodyDecode    = errors.New("body decode failed")
	ErrBodyEncode    = errors.New("body encode failed")
	ErrPathParameter = errors.New("path parameter error")
	ErrNotAllowed    = errors.New("method not allowed")
)
