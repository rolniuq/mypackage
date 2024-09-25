package validator

import (
	"fmt"
	"reflect"
)

type ValidateFunc func() error

type Personal struct {
	Name string
	Age  string
}

type Validator[T any] struct {
	m    map[string]ValidateFunc
	data T
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

func (v *Validator[T]) IsFieldValid(field string) bool {
	arr := reflect.ValueOf(v.data).Elem()
	for i := 0; i < arr.NumField(); i++ {
		if arr.Type().Field(i).Name == field {
			return true
		}
	}

	return false
}

func (v *Validator[T]) Register(name string, fn ValidateFunc) *Validator[T] {
	v.m[name] = fn
	return v
}

func (v *Validator[T]) Validate() error {
	for name, fn := range v.m {
		if !v.IsFieldValid(name) {
			continue
		}

		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func Sample() error {
	p := &Personal{
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
