package controller

import "errors"

var (
	ErrMissingParameter = errors.New("missing parameter")
	ErrInvalidParameter = errors.New("invalid parameter")
)
