package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*SecurityRequirements)(nil)

// https://spec.openapis.org/oas/v3.1.0#security-requirement-object
type SecurityRequirements []map[string][]string

func (s *SecurityRequirements) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_SECURITY_KEY)
	if node != nil {
		err := json.Unmarshal(*node, s)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Security list: %v", err)
		}
	}
	return nil
}
