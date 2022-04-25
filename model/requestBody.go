package model

// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`

	Description string     `json:"description,omitempty" diff:"description"`
	Content     ContentMap `json:"content,omitempty" diff:"content"`
	Required    bool       `json:"required,omitempty" diff:"required"`
}
