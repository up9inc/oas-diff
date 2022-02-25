package model

import (
	file "github.com/up9inc/oas-diff/json"
)

type Model interface {
	// Each model struct must have its own parse logic
	Parse(file file.JsonFile) error
}

// TODO: Remove omitempty from required properties related to OAS 3.1
// TODO: Support Extensions
// TODO: Numbers should be uint64 or int/uint32?
type Schema struct {
	// Schema
	OneOf      []*Schema          `json:"oneOf,omitempty" diff:"oneOf"`
	AnyOf      []*Schema          `json:"anyOf,omitempty" diff:"anyOf"`
	AllOf      []*Schema          `json:"allOf,omitempty" diff:"allOf"`
	Not        *Schema            `json:"not,omitempty" diff:"not"`
	Properties map[string]*Schema `json:"properties,omitempty" diff:"properties"`
	Items      interface{}        `json:"items,omitempty" diff:"items"` // nil or *Schema or []*Schema
	Enum       []interface{}      `json:"enum,omitempty" diff:"enum"`
	Default    interface{}        `json:"default,omitempty" diff:"default"`

	// Bool
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Deprecated      bool `json:"deprecated,omitempty" diff:"deprecated"`

	// String
	Comment     string   `json:"$comment,omitempty" diff:"$comment"`
	Type        string   `json:"type,omitempty" diff:"type"`
	Title       string   `json:"title,omitempty" diff:"title"`
	Format      string   `json:"format,omitempty" diff:"format"`
	Description string   `json:"description,omitempty" diff:"description"`
	Pattern     string   `json:"pattern,omitempty" diff:"pattern"`
	Required    []string `json:"required,omitempty" diff:"required"`

	// Int
	MinItems uint64 `json:"minItems,omitempty" diff:"minItems"`

	// Examples/docs
	Example      interface{}   `json:"example,omitempty" diff:"example"`
	Examples     []interface{} `json:"examples,omitempty" diff:"examples"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" diff:"externalDocs"`
}

// TODO: externalDocs should be a $ref string?
type ExternalDocs struct {
	Description string `json:"description,omitempty" diff:"description"`
	URL         string `json:"url,omitempty" diff:"url"`
}

type SecurityRequirements []map[string][]string
