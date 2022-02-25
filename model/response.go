package model

type Header struct {
	Parameter
}

type Response struct {
	Ref         string             `json:"$ref,omitempty" diff:"$ref"`
	Description string             `json:"description,omitempty" diff:"description"`
	Schema      *Schema            `json:"schema,omitempty" diff:"schema"`
	Headers     map[string]*Header `json:"headers,omitempty" diff:"headers"`
	Examples
}
