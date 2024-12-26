package abstractions

import "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"

type IUserRepository interface {
	Create(email string, password string) error

	GetByEmail(email string) (*models.User, error)
}
