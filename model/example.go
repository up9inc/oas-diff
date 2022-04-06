package model

type ExamplesInterface interface {
	IgnoreExamples()
}

type ExamplesMap map[string]*Example

// https://spec.openapis.org/oas/v3.1.0#example-object
type Example struct {
	Summary       string      `json:"summary,omitempty" diff:"summary"`
	Description   string      `json:"description,omitempty" diff:"description"`
	Value         interface{} `json:"value,omitempty" diff:"value"`
	ExternalValue string      `json:"externalValue,omitempty" diff:"externalValue"`
}
