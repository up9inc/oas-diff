package model

type Headers map[string]*Header

// https://spec.openapis.org/oas/v3.1.0#header-object
type Header struct {
	Description     string                 `json:"description,omitempty" diff:"description"`
	Required        bool                   `json:"required,omitempty" diff:"required"`
	Deprecated      bool                   `json:"deprecated,omitempty" diff:"deprecated"`
	AllowEmptyValue bool                   `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Style           string                 `json:"style,omitempty" diff:"style"`
	Explode         bool                   `json:"explode,omitempty" diff:"explode"`
	AllowReserved   bool                   `json:"allowReserved,omitempty" diff:"allowReserved"`
	Schema          *Schema                `json:"schema,omitempty" diff:"schema"`
	Example         interface{}            `json:"example,omitempty" diff:"example"`
	Examples        map[string]interface{} `json:"examples,omitempty" diff:"examples"`
	Content         Content                `json:"content,omitempty" diff:"content"`
}
