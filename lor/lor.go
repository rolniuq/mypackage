package lor

import (
	"encoding/json"
	"os"
)

func Map[T any, R any](arr []T, f func(T) R) []R {
	result := make([]R, len(arr), cap(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func Filter[T any](arr []T, f func(T) bool) []T {
	result := make([]T, 0, len(arr))

	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[T any, R any](arr []T, init R, f func(R, T) R) R {
	for _, v := range arr {
		init = f(init, v)
	}

	return init
}

func WriteJsonFile[T any](t *T, path string) error {
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

func StructToStruct[T, U any](t *T) (*U, error) {
	var u U
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
