package model

// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Description string  `json:"description,omitempty" diff:"description"`
	Content     Content `json:"content,omitempty" diff:"content"`
	Required    bool    `json:"required,omitempty" diff:"required"`
}
