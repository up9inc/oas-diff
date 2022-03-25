package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*ExternalDoc)(nil)

// https://spec.openapis.org/oas/v3.1.0#external-documentation-object
type ExternalDoc struct {
	Description string `json:"description,omitempty" diff:"description"`
	URL         string `json:"url,omitempty" diff:"url"`
}

func (e *ExternalDoc) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_EXTERNAL_DOCS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, e)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal ExternalDoc struct: %v", err)
		}
	}
	return nil
}
