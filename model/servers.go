package model

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
