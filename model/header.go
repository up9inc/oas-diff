package model

type Headers map[string]Header

// https://spec.openapis.org/oas/v3.1.0#header-object
type Header struct {
	Description      string                 `json:"description,omitempty" diff:"description"`
	CollectionFormat string                 `json:"collectionFormat,omitempty" diff:"collectionFormat"`
	Type             string                 `json:"type,omitempty" diff:"type"`
	Style            string                 `json:"style,omitempty" diff:"style"`
	Explode          bool                   `json:"explode,omitempty" diff:"explode"`
	AllowReserved    bool                   `json:"allowReserved,omitempty" diff:"allowReserved"`
	Format           string                 `json:"format,omitempty" diff:"format"`
	Pattern          string                 `json:"pattern,omitempty" diff:"pattern"`
	AllowEmptyValue  bool                   `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Required         bool                   `json:"required,omitempty" diff:"required"`
	Deprecated       bool                   `json:"deprecated,omitempty" diff:"deprecated"`
	UniqueItems      bool                   `json:"uniqueItems,omitempty" diff:"uniqueItems"`
	ExclusiveMin     bool                   `json:"exclusiveMinimum,omitempty" diff:"exclusiveMinimum"`
	ExclusiveMax     bool                   `json:"exclusiveMaximum,omitempty" diff:"exclusiveMaximum"`
	Schema           Schema                 `json:"schema,omitempty" diff:"schema"`
	Items            Schema                 `json:"items,omitempty" diff:"items"`
	Content          Content                `json:"content,omitempty" diff:"content"`
	Enum             []interface{}          `json:"enum,omitempty" diff:"enum"`
	MultipleOf       float64                `json:"multipleOf,omitempty" diff:"multipleOf"`
	Minimum          float64                `json:"minimum,omitempty" diff:"minimum"`
	Maximum          float64                `json:"maximum,omitempty" diff:"maximum"`
	MaxLength        uint64                 `json:"maxLength,omitempty" diff:"maxLength"`
	MaxItems         uint64                 `json:"maxItems,omitempty" diff:"maxItems"`
	MinLength        uint64                 `json:"minLength,omitempty" diff:"minLength"`
	MinItems         uint64                 `json:"minItems,omitempty" diff:"minItems"`
	Default          interface{}            `json:"default,omitempty" diff:"default"`
	Example          interface{}            `json:"example,omitempty" diff:"example"`
	Examples         map[string]interface{} `json:"examples,omitempty" diff:"examples"`
}
