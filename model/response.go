package model

type Response struct {
	Ref         string        `json:"$ref,omitempty" diff:"$ref"`
	Description string        `json:"description,omitempty" diff:"description"`
	Schema      *Schema       `json:"schema,omitempty" diff:"schema"`
	Headers     Headers       `json:"headers,omitempty" diff:"headers"`
	Example     interface{}   `json:"example,omitempty" diff:"example"`
	Examples    []interface{} `json:"examples,omitempty" diff:"examples"`
}
