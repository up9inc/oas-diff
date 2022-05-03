package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Servers)(nil)

// make sure we implement the Array interface
var _ Array = (*Servers)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*Server)(nil)

// make sure we implement the Descriptions interface
var _ DescriptionsInterface = (*ServerVariable)(nil)

type Servers []*Server
type ServerVariablesMap map[string]*ServerVariable

// https://spec.openapis.org/oas/v3.1.0#server-object
type Server struct {
	URL         string             `json:"url,omitempty" diff:"url,identifier"`
	Description string             `json:"description,omitempty" diff:"description"`
	Variables   ServerVariablesMap `json:"variables,omitempty" diff:"variables"`
}

func (s *Server) IgnoreDescriptions() {
	if s != nil && len(s.Description) > 0 {
		s.Description = ""
	}
}

// https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" diff:"enum"`
	Default     string   `json:"default,omitempty" diff:"default"`
	Description string   `json:"description,omitempty" diff:"description"`
}

func (s *ServerVariable) IgnoreDescriptions() {
	if s != nil && len(s.Description) > 0 {
		s.Description = ""
	}
}

func (s *Servers) Parse(file file.JsonFile) error {
	node := file.GetNodeData(OAS_SERVERS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, s)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal Servers struct: %v", err)
		}
	}

	return nil
}

func (s Servers) GetName() string {
	return OAS_SERVERS_KEY
}

func (s Servers) GetIdentifierName() string {
	return "url"
}

func (s Servers) SearchByIdentifier(identifier interface{}) (int, error) {
	url, ok := identifier.(string)
	if !ok {
		return -1, fmt.Errorf("invalid identifier for %s model, must be a string", s.GetName())
	}

	for k, v := range s {
		if v.URL == url {
			return k, nil
		}
	}

	return -1, nil
}

func (s Servers) FilterIdentifiers() []*ArrayIdentifierFilter {
	var result []*ArrayIdentifierFilter
	for i, d := range s {
		if len(d.URL) > 0 {
			result = append(result, &ArrayIdentifierFilter{
				Name:  d.URL,
				Index: i,
			})
		}
	}
	return result
}
