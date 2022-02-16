package model

import (
	"encoding/json"
	"errors"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

// make sure this model implements the Array interface
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

func ParseServers(file file.JsonFile) (*Servers, error) {
	var serversModel Servers
	node := file.GetNodeData(OAS_SERVERS_KEY)
	if node != nil {
		err := json.Unmarshal(*node, &serversModel)
		if err != nil {
			return nil, fmt.Errorf("failed to Unmarshal Servers struct: %v", err)
		}
	}

	return &serversModel, nil
}

func (s Servers) SearchByIdentifier(identifier interface{}) (int, interface{}, error) {
	url, ok := identifier.(string)
	if !ok {
		return -1, nil, errors.New("invalid identifier for servers model, must be a string")
	}

	for k, v := range s {
		if v.URL == url {
			return k, v, nil
		}
	}

	return -1, nil, nil
}
