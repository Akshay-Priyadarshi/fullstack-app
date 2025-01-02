package repositories

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/models"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func handleError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		if strings.Contains(pgErr.Error(), "SQLSTATE 23505") {
			return ErrRepositoryNonUnique
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		return ErrRepositoryNoRows
	}
	return err
}

type UserRepository struct{}

func (u *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
	`
	_, err := server.AppServer.DB.Exec(query, user.Id, user.Email, user.Password)
	if err != nil {
		return handleError(err)
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
		return nil, handleError(err)
	}
	return &user, nil
}

func (ur *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE id = $1
	`
	var user models.User
	err := server.AppServer.DB.Get(&user, query, id)
	if err != nil {
		return nil, handleError(err)
	}
	return &user, nil
}

func (ur *UserRepository) UpdatePassword(id uuid.UUID, passwordHash string) error {
	query := `
		UPDATE users
		SET password = $1
		WHERE id = $2
	`
	_, err := server.AppServer.DB.Exec(query, passwordHash, id)
	if err != nil {
		return handleError(err)
	}
	return nil
}
