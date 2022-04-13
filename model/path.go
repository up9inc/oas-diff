package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*PathsMap)(nil)

type PathsMap map[string]*PathItem

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

func (p *PathsMap) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_PATHS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, p)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal PathsMap: %v", err)
		}
	}

	return nil
}

func (p *PathItem) Parse(file file.JsonFile, path string) error {
	node := file.GetNodeData(fmt.Sprintf("%s.%s", OAS_PATHS_KEY, path))
	if node != nil {
		err := json.Unmarshal(*node, p)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal PathItem struct: %v", err)
		}
	}

	return nil
}

func (p *PathItem) ParseFromNode(node *[]byte) error {
	if node != nil {
		err := json.Unmarshal(*node, p)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal PathItem struct: %v", err)
		}
	}

	return nil
}

func (p PathItem) GetOperationsName() []string {
	var operations []string

	if p.Connect != nil {
		operations = append(operations, "connect")
	} else if p.Delete != nil {
		operations = append(operations, "delete")
	} else if p.Get != nil {
		operations = append(operations, "get")
	} else if p.Head != nil {
		operations = append(operations, "head")
	} else if p.Options != nil {
		operations = append(operations, "options")
	} else if p.Patch != nil {
		operations = append(operations, "patch")
	} else if p.Post != nil {
		operations = append(operations, "post")
	} else if p.Put != nil {
		operations = append(operations, "put")
	} else if p.Trace != nil {
		operations = append(operations, "trace")
	}

	return operations
}
