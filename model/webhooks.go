package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*WebhooksMap)(nil)

// Map[string, Path Item Object | Reference Object] ]
// https://spec.openapis.org/oas/v3.1.0#schema
type WebhooksMap map[string]*PathItem

func (w *WebhooksMap) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_WEBHOOKS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, w)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Webhooks Map: %v", err)
		}
	}

	return nil
}
