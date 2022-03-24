package model

type ResponsesMap map[string]*Response
type LinksMap map[string]*Link

// https://spec.openapis.org/oas/v3.1.0#response-object
type Response struct {
	Description string     `json:"description,omitempty" diff:"description"`
	Headers     HeadersMap `json:"headers,omitempty" diff:"headers"`
	Content     ContentMap `json:"content,omitempty" diff:"content"`
	Links       LinksMap   `json:"links,omitempty" diff:"links"`
}

// https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	OperationRef string                 `json:"operationRef,omitempty" diff:"operationRef"`
	OperationID  string                 `json:"operationId,omitempty" diff:"operationId"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" diff:"parameters"`
	RequestBody  interface{}            `json:"requestBody,omitempty" diff:"requestBody"`
	Description  string                 `json:"description,omitempty" diff:"description"`
	Server       *Server                `json:"server,omitempty" diff:"server"`
}
