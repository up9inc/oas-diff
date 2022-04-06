package model

import (
	file "github.com/up9inc/oas-diff/json"
)

type Model interface {
	// Each model struct must have its own parse logic
	Parse(file file.JsonFile) error
}

type AnyMap map[string]interface{}
type AnySlice []interface{}
type StringsMap map[string]string

type MapsConstraint[V any] interface {
	~map[string]V
}

type SlicesConstraint[V any] interface {
	~[]V
}

type ExtensionProps struct {
	Extensions map[string]interface{} `json:"-" diff:"-"`
}
