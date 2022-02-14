package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

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

// It should only return 1 result, the url identifier is unique
func (s Servers) FilterByURL(url string) (int, *Server) {
	for k, v := range s {
		if v.URL == url {
			return k, v
		}
	}

	return -1, nil
}
