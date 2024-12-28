package repositoryerrors

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrRepositoryDuplicateValue = errors.New("duplicate value")
	ErrRepositoryNoRows         = errors.New("no rows")
	ErrRepositoryTooManyRows    = errors.New("too many rows")
)

type RepositoryError error

func Mapper(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		if strings.Contains(pgErr.Error(), "SQLSTATE 23505") {
			return ErrRepositoryDuplicateValue
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		return ErrRepositoryNoRows
	}
	return err
}

func New(err error) RepositoryError {
	return RepositoryError(fmt.Errorf("repository error: %w", Mapper(err)))
}
