package passwords

import "fmt"

type PasswordError error

func NewPasswordError(err error) PasswordError {
	return PasswordError(fmt.Errorf("password error: %w", err))
}
