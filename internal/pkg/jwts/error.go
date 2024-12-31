package jwts

import (
	"errors"
)

var (
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrTokenParsingFailed   = errors.New("token parsing failed")
	ErrSigningFailed        = errors.New("signing failed")
)
