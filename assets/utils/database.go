package asset_utils

import (
	"database/sql"

	io "github.com/TimTwigg/Manwe/utils/io"
	logger "github.com/TimTwigg/Manwe/utils/log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// Retrieve the database URL from the environment variables
func GetDBURL() (string, error) {
	env, err := io.GetEnv()
	if err != nil {
		return "", err
	}
	dburl := env["DATABASE_URL"]
	if dburl == "" {
		return "", nil
	}
	return dburl, nil
}

// Open a connection to the database
func GetDB() (*sql.DB, error) {
	dburl, err := GetDBURL()
	if err != nil {
		return nil, err
	}
	if dburl == "" {
		return nil, nil
	}

	db, err := sql.Open("postgres", dburl)
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
	logger.Info("Executing SQL: ", query, args)
	result, err := db.Exec(query, args...)
	logger.Info("SQL Execution Result: ", result)
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
