package model

// https://spec.openapis.org/oas/v3.1.0#operation-object
type Operation struct {
	Tags         []string             `json:"tags,omitempty" diff:"tags"`
	Summary      string               `json:"summary,omitempty" diff:"summary"`
	Description  string               `json:"description,omitempty" diff:"description"`
	ExternalDocs *ExternalDoc         `json:"externalDocs,omitempty" diff:"externalDocs"`
	OperationID  string               `json:"operationId,omitempty" diff:"operationId"`
	Parameters   Parameters           `json:"parameters,omitempty" diff:"parameters"`
	RequestBody  *RequestBody         `json:"requestBody,omitempty" diff:"requestBody"`
	Responses    ResponsesMap         `json:"responses,omitempty" diff:"responses"`
	Callbacks    CallbacksMap         `json:"callbacks,omitempty" diff:"callbacks"`
	Deprecated   bool                 `json:"deprecated,omitempty" diff:"deprecated"`
	Consumes     []string             `json:"consumes,omitempty" diff:"consumes"`
	Security     SecurityRequirements `json:"security,omitempty" diff:"security"`
	Servers      Servers              `json:"servers,omitempty" diff:"servers"`
}
