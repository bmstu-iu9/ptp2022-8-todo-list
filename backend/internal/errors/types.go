package errors

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrDb            = errors.New("database error")
	ErrValidation    = errors.New("validation failed")
	ErrWrongPassword = errors.New("wrong password")
	ErrBodyDecode    = errors.New("body decode failed")
	ErrBodyEncode    = errors.New("body encode failed")
	ErrPathParameter = errors.New("path parameter error")
)
