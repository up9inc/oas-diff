package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Info)(nil)

type Info struct {
	Title          string   `json:"title" diff:"title"`
	Description    string   `json:"description,omitempty" diff:"description"`
	TermsOfService string   `json:"termsOfService,omitempty" diff:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" diff:"contact"`
	License        *License `json:"license,omitempty" diff:"license"`
	Version        string   `json:"version" diff:"version"`
}

type Contact struct {
	Name  string `json:"name,omitempty" diff:"name"`
	URL   string `json:"url,omitempty" diff:"url"`
	Email string `json:"email,omitempty" diff:"email"`
}

type License struct {
	Name string `json:"name" diff:"name"`
	URL  string `json:"url,omitempty" diff:"url"`
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
