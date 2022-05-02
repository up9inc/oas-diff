package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Tags)(nil)

// make sure we implement the Array interface
var _ Array = (*Tags)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Tag)(nil)

type Tags []*Tag

// https://spec.openapis.org/oas/v3.1.0#tag-object
type Tag struct {
	Name         string       `json:"name,omitempty" diff:"name,identifier"`
	Description  string       `json:"description,omitempty" diff:"description"`
	ExternalDocs *ExternalDoc `json:"externalDocs,omitempty" diff:"externalDocs"`
}

func (t *Tags) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_TAGS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, t)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Tags array: %v", err)
		}
	}
	return nil
}

func (t Tags) GetName() string {
	return OAS_TAGS_KEY
}

func (t Tags) GetIdentifierName() string {
	return "name"
}

func (t Tags) SearchByIdentifier(identifier interface{}) (int, error) {
	name, ok := identifier.(string)
	if !ok {
		return -1, fmt.Errorf("invalid identifier for %s model, must be a string", t.GetName())
	}

	for k, v := range t {
		if v.Name == name {
			return k, nil
		}
	}

	return -1, nil
}

func (t Tags) FilterIdentifiers() []*ArrayIdentifierFilter {
	var result []*ArrayIdentifierFilter
	for i, d := range t {
		if len(d.Name) > 0 {
			result = append(result, &ArrayIdentifierFilter{
				Name:  d.Name,
				Index: i,
			})
		}
	}
	return result
}

func (t *Tag) IgnoreDescriptions() {
	if t != nil && len(t.Description) > 0 {
		t.Description = ""
	}
}
