package models

import (
	"errors"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/api/services"
	"github.com/google/uuid"
)

type User struct {
	Identity[uuid.UUID]
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) UpdateEmail(updatedEmail string) {
	u.Email = updatedEmail
}

func (u *User) UpdatePassword(oldPassword string, newPassword string) error {
	if err := services.ComparePassword(oldPassword, u.Password); err != nil {
		return errors.New("old password is incorrect")
	}
	newHashedPassword, err := services.HashPassword(newPassword)
	if err != nil {
		return err
	}
	u.Password = newHashedPassword
	return nil
}
