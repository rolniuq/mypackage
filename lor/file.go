package lor

import (
	"encoding/json"
	"os"
)

func WriteJsonFile[T any](path string, t T) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0644)
}

func ReadJsonFile[T any](path string) (*T, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var t T
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
