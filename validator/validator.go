package validator

import (
	"fmt"
	"reflect"
)

type ValidateFunc func() error

type Validator[T any] struct {
	m    map[string]ValidateFunc
	data T
}

func (v *Validator[T]) isFieldValid(field string) bool {
	arr := reflect.ValueOf(v.data).Elem()
	for i := 0; i < arr.NumField(); i++ {
		if arr.Type().Field(i).Name == field {
			return true
		}
	}

	return false
}

func NewValidator[T any](data T) *Validator[T] {
	return &Validator[T]{
		m:    make(map[string]ValidateFunc),
		data: data,
	}
}

func (v *Validator[T]) Reset() {
	v.m = make(map[string]ValidateFunc)
}

func (v *Validator[T]) Register(name string, fn ValidateFunc) *Validator[T] {
	v.m[name] = fn
	return v
}

func (v *Validator[T]) Validate() error {
	for name, fn := range v.m {
		if !v.isFieldValid(name) {
			continue
		}

		if err := fn(); err != nil {
			return err
		}
	}

	v.Reset()

	return nil
}

func Sample() error {
	type personal struct {
		Name string
		Age  string
	}

	p := &personal{
		Name: "foo",
		Age:  "10",
	}

	v := NewValidator(p)
	v.Register("Name", func() error {
		if p.Name != "" {
			return fmt.Errorf("name is empty")
		}

		return nil
	}).Register("Age", func() error {
		return nil
	}).Register("Hello", func() error {
		return nil
	})

	err := v.Validate()
	if err != nil {
		return err
	}

	return nil
}
