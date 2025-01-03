package abstractions

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(user *models.User) error

	GetByEmail(email string) (*models.User, error)

	GetById(id uuid.UUID) (*models.User, error)

	UpdatePassword(id uuid.UUID, passwordHash string) error
}
