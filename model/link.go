package model

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Response)(nil)

type LinksMap map[string]*Link

// https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	OperationRef string      `json:"operationRef,omitempty" diff:"operationRef"`
	OperationID  string      `json:"operationId,omitempty" diff:"operationId"`
	Parameters   AnyMap      `json:"parameters,omitempty" diff:"parameters"`
	RequestBody  interface{} `json:"requestBody,omitempty" diff:"requestBody"`
	Description  string      `json:"description,omitempty" diff:"description"`
	Server       *Server     `json:"server,omitempty" diff:"server"`

	// Reference Object
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`
}

func (l *Link) IgnoreDescriptions() {
	if l != nil && len(l.Description) > 0 {
		l.Description = ""
	}
}
