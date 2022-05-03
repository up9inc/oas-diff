package model

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*RequestBody)(nil)

type RequestBodiesMap map[string]*RequestBody

// https://spec.openapis.org/oas/v3.1.0#request-body-object
type RequestBody struct {
	Description string     `json:"description,omitempty" diff:"description"`
	Content     ContentMap `json:"content,omitempty" diff:"content"`
	Required    bool       `json:"required,omitempty" diff:"required"`

	// Reference object
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`
}

func (r *RequestBody) IgnoreDescriptions() {
	if r != nil && len(r.Description) > 0 {
		r.Description = ""
	}
}
