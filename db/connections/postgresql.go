package connections

import (
	"log"

	errorhandling "github.com/Akshay-Priyadarshi/fullstack-app/internal/pkg/error_handling"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitializeDB(POSTGRES_URL string) *sqlx.DB {
	DB := errorhandling.PanicIfErrorOrResponse(sqlx.Open("pgx", POSTGRES_URL))
	errorhandling.PanicIfError(DB.Ping())
	log.Println("Connected to PostgreSQL!")
	return DB
}
