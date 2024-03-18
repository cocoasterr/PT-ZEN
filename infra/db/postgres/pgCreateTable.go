package PGConfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func CreateDatabase(dbName string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to connect .env")
		return "", err
	}
	db_dsn := os.Getenv("PG_DB_DSN")
	db_dsn_con := fmt.Sprintf("%s?sslmode=disable", db_dsn)
	db, err := sql.Open("postgres", db_dsn_con)

	if err != nil {
		log.Fatal(err)
	}
	createDatabaseQuery := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s'", dbName)

	var existingDatabaseName string
	if err = db.QueryRow(createDatabaseQuery).Scan(&existingDatabaseName); err != nil {
		if err == sql.ErrNoRows {
			createDatabaseQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
			_, err := db.Exec(createDatabaseQuery)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	new_db_dsn := fmt.Sprintf("%s%s?sslmode=disable", db_dsn, dbName)
	log.Println("Database Created!")
	return new_db_dsn, nil
}

func CreateTable(db *sql.DB) error {
	sqlScript := `
	-- Membuat tabel users
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		firstname VARCHAR,
		lastname VARCHAR,
		phone_number VARCHAR,
		password VARCHAR(255) NOT NULL,
		deleted_at TIMESTAMP
	);

	-- Membuat tabel product
	CREATE TABLE IF NOT EXISTS product (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255),
		harga INT,
		stock INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_by VARCHAR(255),
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_by VARCHAR(255),
		deleted_at TIMESTAMP
	);

	-- Membuat tabel pesanan
	CREATE TABLE IF NOT EXISTS pesanan (
		id VARCHAR(255) PRIMARY KEY,
		product_id VARCHAR(255),
		user_id VARCHAR(255),
		qty INT,
		status BOOLEAN,
		deleted_at TIMESTAMP,
		CONSTRAINT pesanan_product_id_fkey FOREIGN KEY (product_id) REFERENCES product(id),
		CONSTRAINT pesanan_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
	);
	
	`

	_, err := db.Exec(sqlScript)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Println("table already exists")
			return nil
		}
	}
	log.Println("Table Created!")
	return nil
}
