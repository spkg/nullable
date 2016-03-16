package nullable

import "errors"

var (
	jsonNull = []byte("null")
)

var (
	errInvalidString  = errors.New("invalid string")
	errInvalidInteger = errors.New("invalid integer")
)
