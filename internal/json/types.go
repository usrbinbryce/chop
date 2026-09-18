package json

import "errors"

var (
	ErrReqTooLarge  = errors.New("request body too large")
	ErrReqBodyEmpty = errors.New("request body is empty")
)
