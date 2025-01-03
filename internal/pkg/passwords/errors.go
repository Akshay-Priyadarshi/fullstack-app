package passwords

import (
	"errors"
)

var (
	ErrPasswordMismatch = errors.New("passwords do not match")
	ErrHashingFailed    = errors.New("password hashing failed")
)
