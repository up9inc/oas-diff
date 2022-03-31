package model

type HeadersMap map[string]*Header

// make sure we implement the Examples interface
var _ ExamplesInterface = (*Header)(nil)

// https://spec.openapis.org/oas/v3.1.0#header-object
type Header struct {
	Description     string      `json:"description,omitempty" diff:"description"`
	Required        bool        `json:"required,omitempty" diff:"required"`
	Deprecated      bool        `json:"deprecated,omitempty" diff:"deprecated"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Style           string      `json:"style,omitempty" diff:"style"`
	Explode         bool        `json:"explode,omitempty" diff:"explode"`
	AllowReserved   bool        `json:"allowReserved,omitempty" diff:"allowReserved"`
	Schema          *Schema     `json:"schema,omitempty" diff:"schema"`
	Example         interface{} `json:"example,omitempty" diff:"example"`
	Examples        AnyMap      `json:"examples,omitempty" diff:"examples"`
	Content         ContentMap  `json:"content,omitempty" diff:"content"`
}

func (h *Header) IgnoreExamples() {
	if h.Example != nil {
		h.Example = nil
	}
	if h.Examples != nil {
		h.Examples = nil
	}
}
