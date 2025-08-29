package errors

import "errors"

var (
	ErrEmptyItem       = errors.New("empty item")
	ErrInvalidItemList = errors.New("invalid item list, must be a number or a comma separated list of numbers")
	ErrInvalidItem     = errors.New("invalid item, must be a number")
)
