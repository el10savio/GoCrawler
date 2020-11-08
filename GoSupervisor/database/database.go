package database

import (
	"database/sql"
	"fmt"
	"os"

	// Postgres DB Driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "crawler"
)

var (
	// Instantiate a shared DB variable
	// to be shared in the package
	db *sql.DB
)

// Initialize attempts to establish a
// connection the Postgres DB
func Initialize() (*sql.DB, error) {
	var err error
	var dbInfo string

	// Initialize DB connection config
	if os.Getenv("POSTGRES_HOST") != "" {
		dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), port, user, password, dbname)
	} else {
		dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}

	// Connect to the postgres DB
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		return db, err
	}

	// Ping the DB
	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

// GetDB return the shared
// *sql.DB vairable
func GetDB() *sql.DB {
	return db
}
