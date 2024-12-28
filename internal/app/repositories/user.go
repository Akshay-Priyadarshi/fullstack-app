package repositories

import (
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	repositoryerrors "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/repositories/repository_errors"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
)

type UserRepository struct{}

func (u *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`
	_, err := server.AppServer.DB.Exec(query, user.Email, user.Password)
	if err != nil {
		return repositoryerrors.New(err)
	}
	return nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`
	var user models.User
	err := server.AppServer.DB.Get(&user, query, email)
	if err != nil {
		return nil, repositoryerrors.New(err)
	}
	return &user, nil
}
