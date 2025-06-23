package asset_utils

import (
	"database/sql"

	io "github.com/TimTwigg/Manwe/utils/io"
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

// Execute a SQL command and return the result
func ExecSQL(db *sql.DB, query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Insert or update a user in the User table
func UpsertUser(db *sql.DB, userID string) error {
	_, err := ExecSQL(db, "INSERT OR IGNORE INTO User (UserName) VALUES (?)", userID)
	if err != nil {
		return err
	}
	return nil
}
