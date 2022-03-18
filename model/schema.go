package model

// TODO: []*Schema should be handled as an array like servers/parameters? If we do, what will be the identifier?
// TODO: Support Extensions
// TODO: Numbers should be uint64 or int/uint32?
// https://spec.openapis.org/oas/v3.1.0#schema-object
type Schema struct {
	// Schema
	OneOf      []Schema          `json:"oneOf,omitempty" diff:"oneOf"`
	AnyOf      []Schema          `json:"anyOf,omitempty" diff:"anyOf"`
	AllOf      []Schema          `json:"allOf,omitempty" diff:"allOf"`
	Not        interface{}       `json:"not,omitempty" diff:"not"`
	Properties map[string]Schema `json:"properties,omitempty" diff:"properties"`
	Items      interface{}       `json:"items,omitempty" diff:"items"` // nil or *Schema or []*Schema
	Enum       []interface{}     `json:"enum,omitempty" diff:"enum"`
	Default    interface{}       `json:"default,omitempty" diff:"default"`

	// Bool
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Deprecated      bool `json:"deprecated,omitempty" diff:"deprecated"`

	// String
	Ref         string   `json:"$ref,omitempty" diff:"$ref"`
	Comment     string   `json:"$comment,omitempty" diff:"$comment"`
	Type        string   `json:"type,omitempty" diff:"type"`
	Title       string   `json:"title,omitempty" diff:"title"`
	Format      string   `json:"format,omitempty" diff:"format"`
	Description string   `json:"description,omitempty" diff:"description"`
	Pattern     string   `json:"pattern,omitempty" diff:"pattern"`
	Required    []string `json:"required,omitempty" diff:"required"`

	// Int
	MinItems uint64 `json:"minItems,omitempty" diff:"minItems"`

	// fixed fields
	Discriminator Discriminator `json:"discriminator,omitempty" diff:"discriminator"`
	XML           interface{}   `json:"xml,omitempty" diff:"xml"`
	ExternalDocs  ExternalDocs  `json:"externalDocs,omitempty" diff:"externalDocs"`
	Example       interface{}   `json:"example,omitempty" diff:"example"`

	Examples []interface{} `json:"examples,omitempty" diff:"examples"`
}

type Discriminator struct {
	PropertyName string            `json:"propertyName" diff:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty" diff:"mapping"`
}
