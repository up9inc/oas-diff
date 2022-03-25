package model

import (
	file "github.com/up9inc/oas-diff/json"
)

type Model interface {
	// Each model struct must have its own parse logic
	Parse(file file.JsonFile) error
}

type AnyMap map[string]interface{}
type StringsMap map[string]string

// https://spec.openapis.org/oas/v3.1.0#security-requirement-object
type SecurityRequirements []map[string][]string
