package PGConfig

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPG(db_dsn string) (*sql.DB, error) {
	db, _ := sql.Open("postgres", db_dsn)

	log.Println("connected to postgres!")
	return db, nil
}
