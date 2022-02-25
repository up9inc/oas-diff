package model

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	file "github.com/up9inc/oas-diff/json"
)

type Model interface {
	// Each model struct must have its own parse logic
	Parse(file file.JsonFile) error
}

// TODO: SchemaRef not working as expected
type SchemaRef struct {
	Ref   string
	Value *jsonschema.Schema
}

type ExternalDocs struct {
	Description string `json:"description,omitempty" diff:"description"`
	URL         string `json:"url,omitempty" diff:"url"`
}

type SecurityRequirements []map[string][]string
