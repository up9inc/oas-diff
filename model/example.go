package model

type ExamplesInterface interface {
	IgnoreExamples()
}

type ExamplesMap map[string]*Example

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Example)(nil)

// https://spec.openapis.org/oas/v3.1.0#example-object
type Example struct {
	Ref string `json:"$ref,omitempty" diff:"$ref"`

	Summary       string      `json:"summary,omitempty" diff:"summary"`
	Description   string      `json:"description,omitempty" diff:"description"`
	Value         interface{} `json:"value,omitempty" diff:"value"`
	ExternalValue string      `json:"externalValue,omitempty" diff:"externalValue"`
}

func (e *Example) IgnoreDescriptions() {
	if e != nil && len(e.Description) > 0 {
		e.Description = ""
	}
}
