package utils

import (
	"encoding/json"
	"os"
)

func ReadFile(address string) (string, error) {
	file_contents, err := os.ReadFile(address)
	if err != nil {
		return "", err
	}
	return string(file_contents), nil
}

func ReadJSON(address string, target any) error {
	file_contents, err := ReadFile(address)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(file_contents), target)
}

func ListDir(address string) ([]string, error) {
	files, err := os.ReadDir("assets/languages")
	if err != nil {
		return nil, err
	}

	file_names := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			file_names = append(file_names, file.Name())
		}
	}

	return file_names, nil
}

func ApplyToAll(address string, function func(string) error) error {
	files, err := ListDir(address)
	if err != nil {
		return err
	}

	for _, file := range files {
		file_contents, err := ReadFile(address + "/" + file)
		if err != nil {
			return err
		}

		err = function(file_contents)
		if err != nil {
			return err
		}
	}

	return nil
}
