package model

type Content map[string]*MediaType

// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Schema   *Schema              `json:"schema,omitempty" diff:"schema"`
	Example  interface{}          `json:"example,omitempty" diff:"example"`
	Examples Examples             `json:"examples,omitempty" diff:"examples"`
	Encoding map[string]*Encoding `json:"encoding,omitempty" diff:"encoding"`
}

// https://spec.openapis.org/oas/v3.1.0#encoding-object
type Encoding struct {
	ContentType   string  `json:"contentType,omitempty" diff:"contentType"`
	Headers       Headers `json:"headers,omitempty" diff:"headers"`
	Style         string  `json:"style,omitempty" diff:"style"`
	Explode       bool    `json:"explode,omitempty" diff:"explode"`
	AllowReserved bool    `json:"allowReserved,omitempty" diff:"allowReserved"`
}
