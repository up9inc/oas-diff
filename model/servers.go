package model

import (
	"encoding/json"
	"errors"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure we implement the Model interface
var _ Model = (*Servers)(nil)

// make sure we implement the Array interface
var _ Array = (*Servers)(nil)

type Servers []*Server

type Server struct {
	URL         string                     `json:"url" diff:"url,identifier"`
	Description string                     `json:"description,omitempty" diff:"description"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" diff:"variables"`
}

type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" diff:"enum"`
	Default     string   `json:"default,omitempty" diff:"default"`
	Description string   `json:"description,omitempty" diff:"description"`
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

func (s Servers) GetIdentifierName() string {
	return "url"
}

func (s Servers) SearchByIdentifier(identifier interface{}) (int, error) {
	url, ok := identifier.(string)
	if !ok {
		return -1, errors.New("invalid identifier for servers model, must be a string")
	}

	for k, v := range s {
		if v.URL == url {
			return k, nil
		}
	}

	return -1, nil
}
