package model

type Contact struct {
	Name  string `json:"name,omitempty" diff:"name"`
	URL   string `json:"url,omitempty" diff:"url"`
	Email string `json:"email,omitempty" diff:"email"`
}

type License struct {
	Name string `json:"name" diff:"name"`
	URL  string `json:"url,omitempty" diff:"url"`
}

type Info struct {
	Title          string   `json:"title" diff:"title"`
	Description    string   `json:"description,omitempty" diff:"description"`
	TermsOfService string   `json:"termsOfService,omitempty" diff:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" diff:"contact"`
	License        *License `json:"license,omitempty" diff:"license"`
	Version        string   `json:"version" diff:"version"`
}

type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" diff:"enum"`
	Default     string   `json:"default,omitempty" diff:"default"`
	Description string   `json:"description,omitempty" diff:"description"`
}

type Server struct {
	URL         string                     `json:"url" diff:"url"`
	Description string                     `json:"description,omitempty" diff:"description"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty" diff:"variables"`
}

type Servers []*Server
