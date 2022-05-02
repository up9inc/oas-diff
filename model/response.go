package model

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Response)(nil)

type ResponsesMap map[string]*Response

// https://spec.openapis.org/oas/v3.1.0#response-object
// TODO: Understand the "default" fixed field
type Response struct {
	Description string     `json:"description,omitempty" diff:"description"`
	Headers     HeadersMap `json:"headers,omitempty" diff:"headers"`
	Content     ContentMap `json:"content,omitempty" diff:"content"`
	Links       LinksMap   `json:"links,omitempty" diff:"links"`

	// Reference object
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`
}

func (r *Response) IgnoreDescriptions() {
	if r != nil && len(r.Description) > 0 {
		r.Description = ""
	}
}
