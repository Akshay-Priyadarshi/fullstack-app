package jwts

import "fmt"

type JwtError error

func NewJwtError(err error) JwtError {
	return JwtError(fmt.Errorf("jwt error: %w", err))
}
