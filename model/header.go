package model

type HeadersMap map[string]*Header

// make sure we implement the Examples interface
var _ ExamplesInterface = (*Header)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Header)(nil)

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

	// Reference object
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`
}

func (h *Header) IgnoreExamples() {
	if h != nil {
		if h.Example != nil {
			h.Example = nil
		}
		if h.Examples != nil {
			h.Examples = nil
		}
	}
}

func (h *Header) IgnoreDescriptions() {
	if h != nil && len(h.Description) > 0 {
		h.Description = ""
	}
}
