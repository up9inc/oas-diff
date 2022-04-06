package model

type ContentMap map[string]*MediaType
type EncodingMap map[string]*Encoding

// make sure we implement the Examples interface
var _ ExamplesInterface = (*MediaType)(nil)

// https://spec.openapis.org/oas/v3.1.0#media-type-object
type MediaType struct {
	Schema   *Schema     `json:"schema,omitempty" diff:"schema"`
	Example  interface{} `json:"example,omitempty" diff:"example"`
	Examples ExamplesMap `json:"examples,omitempty" diff:"examples"`
	Encoding EncodingMap `json:"encoding,omitempty" diff:"encoding"`
}

// https://spec.openapis.org/oas/v3.1.0#encoding-object
type Encoding struct {
	ContentType   string     `json:"contentType,omitempty" diff:"contentType"`
	Headers       HeadersMap `json:"headers,omitempty" diff:"headers"`
	Style         string     `json:"style,omitempty" diff:"style"`
	Explode       bool       `json:"explode,omitempty" diff:"explode"`
	AllowReserved bool       `json:"allowReserved,omitempty" diff:"allowReserved"`
}

func (m *MediaType) IgnoreExamples() {
	if m.Example != nil {
		m.Example = nil
	}
	if m.Examples != nil {
		m.Examples = nil
	}
}
