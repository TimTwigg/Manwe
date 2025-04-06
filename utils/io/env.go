package io

import (
	godotenv "github.com/joho/godotenv"
)

func GetEnv() (map[string]string, error) {
	env, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return env, nil
}
