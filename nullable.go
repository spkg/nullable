package nullable

//go:generate go run ./cmd/nullable-generate/main.go

import "errors"

var (
	jsonNull = []byte("null")
)

var (
	errInvalidString  = errors.New("invalid string")
	errInvalidInteger = errors.New("invalid integer")
)
