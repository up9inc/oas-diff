package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Components)(nil)

type ParametersMap map[string]*Parameter

// make sure we implement the Examples interface
var _ ExamplesInterface = (*Components)(nil)

// https://spec.openapis.org/oas/v3.1.0#components-object
type Components struct {
	Schemas         SchemasMap         `json:"schemas,omitempty" diff:"schemas"`
	Responses       ResponsesMap       `json:"responses,omitempty" diff:"responses"`
	Parameters      ParametersMap      `json:"parameters,omitempty" diff:"parameters"`
	Examples        ExamplesMap        `json:"examples,omitempty" diff:"examples"`
	RequestBodies   RequestBodiesMap   `json:"requestBodies,omitempty" diff:"requestBodies"`
	Headers         HeadersMap         `json:"headers,omitempty" diff:"headers"`
	SecuritySchemes SecuritySchemesMap `json:"securitySchemes,omitempty" diff:"securitySchemes"`
	Links           LinksMap           `json:"links,omitempty" diff:"links"`
	Callbacks       CallbacksMap       `json:"callbacks,omitempty" diff:"callbacks"`
	PathItems       PathsMap           `json:"pathItems,omitempty" diff:"pathItems"`
}

func (c *Components) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_COMPONENTS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, c)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Components struct: %v", err)
		}
	}

	return nil
}

func (c *Components) IgnoreExamples() {
	if c.Examples != nil {
		c.Examples = nil
	}
}
