package model

// TODO: Add Links?
type Response struct {
	Description string  `json:"description,omitempty" diff:"description"`
	Headers     Headers `json:"headers,omitempty" diff:"headers"`
	Content     Content `json:"content,omitempty" diff:"content"`
}
