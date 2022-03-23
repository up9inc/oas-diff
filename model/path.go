package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Paths)(nil)

type Paths map[string]*PathItem

// https://spec.openapis.org/oas/v3.1.0#path-item-object
type PathItem struct {
	Ref         string     `json:"$ref,omitempty" diff:"$ref"`
	Summary     string     `json:"summary,omitempty" diff:"summary"`
	Description string     `json:"description,omitempty" diff:"description"`
	Connect     *Operation `json:"connect,omitempty" diff:"connect"`
	Delete      *Operation `json:"delete,omitempty" diff:"delete"`
	Get         *Operation `json:"get,omitempty" diff:"get"`
	Head        *Operation `json:"head,omitempty" diff:"head"`
	Options     *Operation `json:"options,omitempty" diff:"options"`
	Patch       *Operation `json:"patch,omitempty" diff:"patch"`
	Post        *Operation `json:"post,omitempty" diff:"post"`
	Put         *Operation `json:"put,omitempty" diff:"put"`
	Trace       *Operation `json:"trace,omitempty" diff:"trace"`
	Servers     Servers    `json:"servers,omitempty" diff:"servers"`
	Parameters  Parameters `json:"parameters,omitempty" diff:"parameters"`
}

func (p *Paths) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_PATHS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, p)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Paths struct: %v", err)
		}
	}

	return nil
}
