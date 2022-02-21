package model

import (
	"encoding/json"
	"errors"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Paths)(nil)

type Paths map[string]*PathItem

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

type Operation struct {
	Summary      string                `json:"summary,omitempty" diff:"summary"`
	Description  string                `json:"description,omitempty" diff:"description"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty" diff:"externalDocs"`
	Tags         []string              `json:"tags,omitempty" diff:"tags"`
	OperationID  string                `json:"operationId,omitempty" diff:"operationId"`
	Parameters   Parameters            `json:"parameters,omitempty" diff:"parameters"`
	Responses    map[string]*Response  `json:"responses" diff:"responses"`
	Consumes     []string              `json:"consumes,omitempty" diff:"consumes"`
	Produces     []string              `json:"produces,omitempty" diff:"produces"`
	Security     *SecurityRequirements `json:"security,omitempty" diff:"security"`
}

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
	Format           string        `json:"format,omitempty" diff:"format"`
	Pattern          string        `json:"pattern,omitempty" diff:"pattern"`
	AllowEmptyValue  bool          `json:"allowEmptyValue,omitempty" diff:"allowEmptyValue"`
	Required         bool          `json:"required,omitempty" diff:"required"`
	UniqueItems      bool          `json:"uniqueItems,omitempty" diff:"uniqueItems"`
	ExclusiveMin     bool          `json:"exclusiveMinimum,omitempty" diff:"exclusiveMinimum"`
	ExclusiveMax     bool          `json:"exclusiveMaximum,omitempty" diff:"exclusiveMaximum"`
	Schema           *SchemaRef    `json:"schema,omitempty" diff:"schema"`
	Items            *SchemaRef    `json:"items,omitempty" diff:"items"`
	Enum             []interface{} `json:"enum,omitempty" diff:"enum"`
	MultipleOf       *float64      `json:"multipleOf,omitempty" diff:"multipleOf"`
	Minimum          *float64      `json:"minimum,omitempty" diff:"minimum"`
	Maximum          *float64      `json:"maximum,omitempty" diff:"maximum"`
	MaxLength        *uint64       `json:"maxLength,omitempty" diff:"maxLength"`
	MaxItems         *uint64       `json:"maxItems,omitempty" diff:"maxItems"`
	MinLength        uint64        `json:"minLength,omitempty" diff:"minLength"`
	MinItems         uint64        `json:"minItems,omitempty" diff:"minItems"`
	Default          interface{}   `json:"default,omitempty" diff:"default"`
}

type Response struct {
	Ref         string             `json:"$ref,omitempty" diff:"$ref"`
	Description string             `json:"description,omitempty" diff:"description"`
	Schema      *SchemaRef         `json:"schema,omitempty" diff:"schema"`
	Headers     map[string]*Header `json:"headers,omitempty" diff:"headers"`
	// TODO: 3.1 examples is an object
	Examples map[string]interface{} `json:"examples,omitempty" diff:"examples"`
}

type Header struct {
	Parameter
}

func (p *Paths) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_PATHS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, p)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Paths struct: %v", err)
		}
	}

	return nil
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
