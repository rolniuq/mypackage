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

func (v *Validator[T]) Validate() []error {
	var errs []error

	for name, fn := range v.m {
		if !v.isFieldValid(name) {
			continue
		}

		if err := fn(); err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func Sample() []error {
	type project struct {
		Name string
	}

	type personal struct {
		Name     string
		Age      string
		Projects []project
	}

	p := &personal{
		Name: "foo",
		Age:  "10",
		Projects: []project{
			{
				Name: "",
			},
		},
	}

	v := NewValidator(p).
		Register("Name", func() error {
			if p.Name != "" {
				return fmt.Errorf("name is empty")
			}

			return nil
		}).
		Register("Age", func() error {
			if p.Age != "10" {
				return fmt.Errorf("age is not 10")
			}

			return nil
		}).
		Register("Hello", func() error {
			return nil
		}).
		Register("Projects", func() error {
			if len(p.Projects) == 0 {
				return fmt.Errorf("projects is empty")
			}

			for _, project := range p.Projects {
				if project.Name == "" {
					return fmt.Errorf("projects name is empty")
				}
			}
			return nil
		})

	errs := v.Validate()
	if errs != nil {
		return errs
	}

	return nil
}
