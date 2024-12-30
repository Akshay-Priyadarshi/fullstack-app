package users

import "fmt"

type ErrUserEmailAlreadyExists struct {
	Email string
}

func (e ErrUserEmailAlreadyExists) Error() string {
	return fmt.Sprintf("user with email '%s' already exists", e.Email)
}

type ErrUserEmailNotFound struct {
	Email string
}

func (e ErrUserEmailNotFound) Error() string {
	return fmt.Sprintf("user with email '%s' not found", e.Email)
}
