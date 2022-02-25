package model

import "errors"

// make sure we implement the Array interface
var _ Array = (*Parameters)(nil)

type Parameters []*Parameter

type Parameter struct {
	Name             string        `json:"name,omitempty" diff:"name,identifier"`
	Ref              string        `json:"$ref,omitempty" diff:"$ref"`
	In               string        `json:"in,omitempty" diff:"in"`
	Description      string        `json:"description,omitempty" diff:"description"`
	CollectionFormat string        `json:"collectionFormat,omitempty" diff:"collectionFormat"`
	Type             string        `json:"type,omitempty" diff:"type"`
	Style            string        `json:"style,omitempty" diff:"style"`
	Format           string        `json:"format,omitempty" diff:"format"`
	Pattern          string        `json:"pattern,omitempty" diff:"pattern"`
	AllowEmptyValue  bool          `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Required         bool          `json:"required,omitempty" diff:"required"`
	UniqueItems      bool          `json:"uniqueItems,omitempty" diff:"uniqueItems"`
	ExclusiveMin     bool          `json:"exclusiveMinimum,omitempty" diff:"exclusiveMinimum"`
	ExclusiveMax     bool          `json:"exclusiveMaximum,omitempty" diff:"exclusiveMaximum"`
	Schema           *Schema       `json:"schema,omitempty" diff:"schema"`
	Items            *Schema       `json:"items,omitempty" diff:"items"`
	Enum             []interface{} `json:"enum,omitempty" diff:"enum"`
	MultipleOf       *float64      `json:"multipleOf,omitempty" diff:"multipleOf"`
	Minimum          *float64      `json:"minimum,omitempty" diff:"minimum"`
	Maximum          *float64      `json:"maximum,omitempty" diff:"maximum"`
	MaxLength        *uint64       `json:"maxLength,omitempty" diff:"maxLength"`
	MaxItems         *uint64       `json:"maxItems,omitempty" diff:"maxItems"`
	MinLength        uint64        `json:"minLength,omitempty" diff:"minLength"`
	MinItems         uint64        `json:"minItems,omitempty" diff:"minItems"`
	Default          interface{}   `json:"default,omitempty" diff:"default"`
	Examples
}

func (p Parameters) SearchByIdentifier(identifier interface{}) (int, error) {
	name, ok := identifier.(string)
	if !ok {
		return -1, errors.New("invalid identifier for parameters model, must be a string")
	}

	for k, v := range p {
		if v.Name == name {
			return k, nil
		}
	}

	return -1, nil
}
