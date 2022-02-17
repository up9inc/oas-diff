package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Paths)(nil)

type Paths map[string]*PathItem

// TODO: Add diff targs
// TODO: Add identifiers to arrays
type PathItem struct {
	Ref         string     `json:"$ref,omitempty"`
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Connect     *Operation `json:"connect,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Head        *Operation `json:"head,omitempty"`
	Options     *Operation `json:"options,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Trace       *Operation `json:"trace,omitempty"`
	Servers     Servers    `json:"servers,omitempty"`
	Parameters  Parameters `json:"parameters,omitempty"`
}

type Operation struct {
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
	Tags         []string              `json:"tags,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Parameters   Parameters            `json:"parameters,omitempty"`
	Responses    map[string]*Response  `json:"responses"`
	Consumes     []string              `json:"consumes,omitempty"`
	Produces     []string              `json:"produces,omitempty"`
	Security     *SecurityRequirements `json:"security,omitempty"`
}

type Parameters []*Parameter

type Parameter struct {
	Ref              string        `json:"$ref,omitempty"`
	In               string        `json:"in,omitempty"`
	Name             string        `json:"name,omitempty"`
	Description      string        `json:"description,omitempty"`
	CollectionFormat string        `json:"collectionFormat,omitempty"`
	Type             string        `json:"type,omitempty"`
	Format           string        `json:"format,omitempty"`
	Pattern          string        `json:"pattern,omitempty"`
	AllowEmptyValue  bool          `json:"allowEmptyValue,omitempty"`
	Required         bool          `json:"required,omitempty"`
	UniqueItems      bool          `json:"uniqueItems,omitempty"`
	ExclusiveMin     bool          `json:"exclusiveMinimum,omitempty"`
	ExclusiveMax     bool          `json:"exclusiveMaximum,omitempty"`
	Schema           *SchemaRef    `json:"schema,omitempty"`
	Items            *SchemaRef    `json:"items,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
	MultipleOf       *float64      `json:"multipleOf,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty"`
	Maximum          *float64      `json:"maximum,omitempty"`
	MaxLength        *uint64       `json:"maxLength,omitempty"`
	MaxItems         *uint64       `json:"maxItems,omitempty"`
	MinLength        uint64        `json:"minLength,omitempty"`
	MinItems         uint64        `json:"minItems,omitempty"`
	Default          interface{}   `json:"default,omitempty"`
}

type Response struct {
	Ref         string             `json:"$ref,omitempty"`
	Description string             `json:"description,omitempty"`
	Schema      *SchemaRef         `json:"schema,omitempty"`
	Headers     map[string]*Header `json:"headers,omitempty"`
	// TODO: 3.1 examples is an object
	Examples map[string]interface{} `json:"examples,omitempty"`
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
