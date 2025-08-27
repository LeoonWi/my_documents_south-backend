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

func New() *sqlx.DB {
	db, err := sqlx.Connect("pgx", "host=localhost port=5433 user=mds_user password=zxcvbn dbname=mds sslmode=disable")
	dieIf(err)

	data, err := os.ReadFile("./schemas/v1/schema.sql")
	dieIf(err)

	schema := string(data)

	db.Exec(schema)
	return db
}
