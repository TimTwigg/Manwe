package io

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(address string) (string, error) {
	file_contents, err := os.ReadFile(address)
	if err != nil {
		return "", err
	}
	return string(file_contents), nil
}

func ReadJSON(address string) (map[string]any, error) {
	file_contents, err := ReadFile(address)
	if err != nil {
		return nil, err
	}
	output := make(map[string]any)
	err = json.Unmarshal([]byte(file_contents), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func ListDir(address string) ([]string, error) {
	files, err := os.ReadDir(address)
	if err != nil {
		return nil, err
	}

	file_names := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			file_names = append(file_names, fmt.Sprintf("%s/%s", address, file.Name()))
		}
	}

	return file_names, nil
}
