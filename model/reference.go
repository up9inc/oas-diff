package model

// https://spec.openapis.org/oas/v3.1.0#reference-object
type Reference struct {
	Ref         string `json:"$ref,omitempty" diff:"$ref"`
	Summary     string `json:"summary,omitempty" diff:"summary"`
	Description string `json:"description,omitempty" diff:"description"`
}
