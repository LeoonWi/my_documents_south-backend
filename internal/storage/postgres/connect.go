package postgres

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func dieIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// New() создает подключение к базе данных и возвращает пул соединений *sqlx.DB
func New() *sqlx.DB {
	cfg := "host=localhost port=5433 user=mds_user password=zxcvbn dbname=mds sslmode=disable"
	db, err := sqlx.Connect("pgx", cfg)
	dieIf(err)

	data, err := os.ReadFile("./schemas/v1/schema.sql")
	dieIf(err)

	schema := string(data)

	if _, err := db.Exec(schema); nil != err {
		dieIf(err)
	}
	return db
}
