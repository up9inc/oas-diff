package model

type RequestBody struct {
	Description string  `json:"description,omitempty" diff:"description"`
	Required    bool    `json:"required,omitempty" diff:"required"`
	Content     Content `json:"content,omitempty" diff:"content"`
}
