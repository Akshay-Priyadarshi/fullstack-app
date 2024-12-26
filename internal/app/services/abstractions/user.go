package abstractions

import "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"

type IUserService interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}
