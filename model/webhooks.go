package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Webhooks)(nil)

// Map[string, Path Item Object | Reference Object] ]
// https://spec.openapis.org/oas/v3.1.0#schema
type Webhooks map[string]*PathItem

func (w *Webhooks) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_WEBHOOKS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, w)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Webhooks struct: %v", err)
		}
	}

	return nil
}
