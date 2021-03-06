package model

import (
	"fmt"
	"strings"
)

// make sure we implement the Array interface
var _ Array = (*Parameters)(nil)

type Parameters []*Parameter

// make sure we implement the Examples interface
var _ ExamplesInterface = (*Parameter)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Parameter)(nil)

// https://spec.openapis.org/oas/v3.1.0#parameter-object
// TODO: Handle Array of Reference Object with our custom array interface?
// TODO: Handle Multiple identifiers when the model is an array of Reference Object
type Parameter struct {
	Name            string      `json:"name,omitempty" diff:"name,identifier"`
	In              string      `json:"in,omitempty" diff:"in"`
	Description     string      `json:"description,omitempty" diff:"description"`
	Required        bool        `json:"required,omitempty" diff:"required"`
	Deprecated      bool        `json:"deprecated,omitempty" diff:"deprecated"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Style           string      `json:"style,omitempty" diff:"style"`
	Explode         bool        `json:"explode,omitempty" diff:"explode"`
	AllowReserved   bool        `json:"allowReserved,omitempty" diff:"allowReserved"`
	Schema          *Schema     `json:"schema,omitempty" diff:"schema"`
	Example         interface{} `json:"example,omitempty" diff:"example"`
	Examples        ExamplesMap `json:"examples,omitempty" diff:"examples"`
	Content         ContentMap  `json:"content,omitempty" diff:"content"`

	// Reference object
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`
}

func (p Parameters) GetName() string {
	return "parameters"
}

func (p Parameters) GetIdentifierName() string {
	return "name"
}

func (p Parameters) SearchByIdentifier(identifier interface{}) (int, error) {
	name, ok := identifier.(string)
	if !ok {
		return -1, fmt.Errorf("invalid identifier for %s model, must be a string", p.GetName())
	}

	for k, v := range p {
		if v.Name == name {
			return k, nil
		}
	}

	return -1, nil
}

func (p Parameters) FilterIdentifiers() []*ArrayIdentifierFilter {
	var result []*ArrayIdentifierFilter
	for i, d := range p {
		if len(d.Name) > 0 {
			result = append(result, &ArrayIdentifierFilter{
				Name:  d.Name,
				Index: i,
			})
		}
	}
	return result
}

func (p Parameter) IsHeader() bool {
	return strings.ToLower(p.In) == "header"
}

func (p Parameter) IsIgnoredWhenLoose() bool {
	name := strings.ToLower(p.Name)
	return strings.HasPrefix(name, "x-") || name == "user-agent"
}

func (p *Parameter) IgnoreExamples() {
	if p.Example != nil {
		p.Example = nil
	}
	if p.Examples != nil {
		p.Examples = nil
	}
}

func (p *Parameter) IgnoreDescriptions() {
	if p != nil && len(p.Description) > 0 {
		p.Description = ""
	}
}
