package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Info)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Info)(nil)

// https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Title          string   `json:"title,omitempty" diff:"title"`
	Summary        string   `json:"summary,omitempty" diff:"summary"`
	Description    string   `json:"description,omitempty" diff:"description"`
	TermsOfService string   `json:"termsOfService,omitempty" diff:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" diff:"contact"`
	License        *License `json:"license,omitempty" diff:"license"`
	Version        string   `json:"version,omitempty" diff:"version"`
}

// https://spec.openapis.org/oas/v3.1.0#contact-object
type Contact struct {
	Name  string `json:"name,omitempty" diff:"name"`
	URL   string `json:"url,omitempty" diff:"url"`
	Email string `json:"email,omitempty" diff:"email"`
}

// https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Name       string `json:"name,omitempty" diff:"name"`
	Identifier string `json:"identifier,omitempty" diff:"identifier"`
	URL        string `json:"url,omitempty" diff:"url"`
}

func (i *Info) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_INFO_KEY)
	if node != nil {
		err := json.Unmarshal(*node, i)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Info struct: %v", err)
		}
	}
	return nil
}

func (i *Info) IgnoreDescriptions() {
	if i != nil && len(i.Description) > 0 {
		i.Description = ""
	}
}
