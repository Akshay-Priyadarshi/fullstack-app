package connections

import (
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitializeDB(POSTGRES_URL string) *sqlx.DB {
	DbPtr, err := sqlx.Open("pgx", POSTGRES_URL)
	if err != nil {
		panic(err)
	}
	err = DbPtr.Ping()
	if err != nil {
		panic(err)
	}
	slog.Info("Connected to PostgreSQL!")
	return DbPtr
}
