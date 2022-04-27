package model

import "fmt"

type ReferencesMap map[string]*Reference
type References []*Reference

// make sure we implement the Array interface
var _ Array = (*References)(nil)

// https://spec.openapis.org/oas/v3.1.0#reference-object
type Reference struct {
	Ref         string `json:"$ref,omitempty" diff:"$ref,identifier"`
	Summary     string `json:"summary,omitempty" diff:"summary"`
	Description string `json:"description,omitempty" diff:"description"`
}

func (r References) GetName() string {
	return "references"
}

func (r References) GetIdentifierName() string {
	return "$ref"
}

func (r References) SearchByIdentifier(identifier interface{}) (int, error) {
	ref, ok := identifier.(string)
	if !ok {
		return -1, fmt.Errorf("invalid identifier for %s model, must be a string", r.GetName())
	}

	for k, v := range r {
		if v.Ref == ref {
			return k, nil
		}
	}

	return -1, nil
}

func (r References) FilterIdentifiers() []*ArrayIdentifierFilter {
	var result []*ArrayIdentifierFilter
	for i, d := range r {
		if len(d.Ref) > 0 {
			result = append(result, &ArrayIdentifierFilter{
				Name:  d.Ref,
				Index: i,
			})
		}
	}
	return result
}
