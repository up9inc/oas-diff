package model

type SchemasMap map[string]*Schema
type SchemasSlice []*Schema

// make sure we implement the Examples interface
var _ ExamplesInterface = (*Schema)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Schema)(nil)

// TODO: []*Schema should be handled as an array like servers/parameters? If we do, what will be the identifier?
// TODO: Support Extensions
// TODO: Numbers should be uint64 or int/uint32?
// https://spec.openapis.org/oas/v3.1.0#schema-object
// https://json-schema.org/draft/2020-12/release-notes.html
type Schema struct {
	// Schema
	Defs                 SchemasMap            `json:"$defs,omitempty" diff:"$defs"`
	OneOf                SchemasSlice          `json:"oneOf,omitempty" diff:"oneOf"`
	AnyOf                SchemasSlice          `json:"anyOf,omitempty" diff:"anyOf"`
	AllOf                SchemasSlice          `json:"allOf,omitempty" diff:"allOf"`
	Not                  *Schema               `json:"not,omitempty" diff:"not"`
	If                   *Schema               `json:"if,omitempty" diff:"if"`
	Then                 *Schema               `json:"then,omitempty" diff:"then"`
	Else                 *Schema               `json:"else,omitempty" diff:"else"`
	Properties           SchemasMap            `json:"properties,omitempty" diff:"properties"`
	PropertyNames        *Schema               `json:"propertyNames,omitempty" diff:"propertyNames"`
	PrefixItems          []*Schema             `json:"prefixItems,omitempty" diff:"prefixItems"`
	Items                *Schema               `json:"items,omitempty" diff:"items"`
	Enum                 []string              `json:"enum,omitempty" diff:"enum"`
	Default              interface{}           `json:"default,omitempty" diff:"default"`
	AdditionalProperties interface{}           `json:"additionalProperties,omitempty" diff:"additionalProperties"`
	Components           map[string]SchemasMap `json:"components,omitempty" diff:"components"`
	Contains             *Schema               `json:"contains,omitempty" diff:"contains"`
	UnevaluatedItems     *Schema               `json:"unevaluatedItems,omitempty" diff:"unevaluatedItems"`
	PatternProperties    SchemasMap            `json:"patternProperties,omitempty" diff:"patternProperties"`
	DependentSchemas     SchemasMap            `json:"dependentSchemas,omitempty" diff:"dependentSchemas"`

	// Bool
	AllowEmptyValue       bool `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Deprecated            bool `json:"deprecated,omitempty" diff:"deprecated"`
	UnevaluatedProperties bool `json:"unevaluatedProperties,omitempty" diff:"unevaluatedProperties"`

	// String
	Ref           string   `json:"$ref,omitempty" diff:"$ref"`
	DynamicAnchor string   `json:"$dynamicAnchor,omitempty" diff:"$dynamicAnchor"`
	DynamicRef    string   `json:"$dynamicRef,omitempty" diff:"$dynamicRef"`
	Comment       string   `json:"$comment,omitempty" diff:"$comment"`
	Type          string   `json:"type,omitempty" diff:"type"`
	Title         string   `json:"title,omitempty" diff:"title"`
	Format        string   `json:"format,omitempty" diff:"format"`
	Description   string   `json:"description,omitempty" diff:"description"`
	Pattern       string   `json:"pattern,omitempty" diff:"pattern"`
	Required      []string `json:"required,omitempty" diff:"required"`

	// Int
	MinProperties    uint64 `json:"minProperties,omitempty" diff:"minProperties"`
	MaxProperties    uint64 `json:"maxProperties,omitempty" diff:"maxProperties"`
	MinItems         uint64 `json:"minItems,omitempty" diff:"minItems"`
	Minimum          uint64 `json:"minimum,omitempty" diff:"minimum"`
	ExclusiveMinimum uint64 `json:"exclusiveMinimum,omitempty" diff:"exclusiveMinimum"`
	Maximum          uint64 `json:"maximum,omitempty" diff:"maximum"`
	ExclusiveMaximum uint64 `json:"exclusiveMaximum,omitempty" diff:"exclusiveMaximum"`
	MultipleOf       uint64 `json:"multipleOf,omitempty" diff:"multipleOf"`

	// fixed fields
	Discriminator *Discriminator `json:"discriminator,omitempty" diff:"discriminator"`
	XML           interface{}    `json:"xml,omitempty" diff:"xml"`
	ExternalDocs  *ExternalDoc   `json:"externalDocs,omitempty" diff:"externalDocs"`
	Example       interface{}    `json:"example,omitempty" diff:"example"`
	Examples      []interface{}  `json:"examples,omitempty" diff:"examples"`

	//Extensions map[string]interface{} `json:"-" diff:"-"`
}

type Discriminator struct {
	PropertyName string     `json:"propertyName,omitempty" diff:"propertyName"`
	Mapping      StringsMap `json:"mapping,omitempty" diff:"mapping"`
}

func (s *Schema) IgnoreExamples() {
	if s.Example != nil {
		s.Example = nil
	}
	if s.Examples != nil {
		s.Examples = nil
	}
}

func (s *Schema) IgnoreDescriptions() {
	if s != nil && len(s.Description) > 0 {
		s.Description = ""
	}
}
