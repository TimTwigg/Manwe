package database

import (
	"database/sql"

	io "github.com/TimTwigg/EncounterManagerBackend/utils/io"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Retrieve the database file path from the environment variables
func GetDBFile() (string, error) {
	env, err := io.GetEnv()
	if err != nil {
		return "", err
	}
	dbfile := env["DBFILE"]
	if dbfile == "" {
		return "", nil
	}
	return dbfile, nil
}

// Open a connection to the database
func GetDB() (*sql.DB, error) {
	dbfile, err := GetDBFile()
	if err != nil {
		return nil, err
	}
	if dbfile == "" {
		return nil, nil
	}

	db, err := sql.Open("sqlite", "file:"+dbfile+"?cache=shared")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Close the database connection
func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

// Execute a SQL Query and return the result
func QuerySQL(db *sql.DB, query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ExecSQL(db *sql.DB, query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
