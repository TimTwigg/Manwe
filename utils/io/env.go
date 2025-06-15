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

func GetEnvVar(key string) (string, error) {
	env, err := GetEnv()
	if err != nil {
		return "", err
	}
	value, exists := env[key]
	if !exists {
		return "", nil
	}
	return value, nil
}
