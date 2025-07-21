package asset_utils

import (
	"context"

	io "github.com/TimTwigg/Manwe/utils/io"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

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
func GetDB() (*pgxpool.Pool, error) {
	dburl, err := GetDBURL()
	if err != nil {
		return nil, err
	}
	if dburl == "" {
		return nil, nil
	}

	pool, err := pgxpool.New(context.Background(), dburl)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}

// Insert or update a user in the User table
func UpsertUser(userID string) error {
	_, err := DBPool.Exec(context.Background(), "INSERT INTO public.user (username) VALUES ($1) ON CONFLICT (username) DO NOTHING", userID)
	if err != nil {
		return err
	}
	return nil
}
