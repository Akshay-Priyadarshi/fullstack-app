package repositories

import (
	"errors"
)

var (
	ErrRepositoryNonUnique   = errors.New("column value must be unique")
	ErrRepositoryNoRows      = errors.New("no rows found")
	ErrRepositoryTooManyRows = errors.New("too many rows found")
)
