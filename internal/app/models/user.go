package models

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/passwords"
	"github.com/google/uuid"
)

type User struct {
	Identity[uuid.UUID]
	Email    string `json:"email"`
	Password string `json:"password"`
	TimeStamps
}

func (u *User) ComparePassword(password string) error {
	if err := passwords.ComparePassword(password, u.Password); err != nil {
		return err
	}
	return nil
}
