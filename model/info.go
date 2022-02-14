package model

import (
	"encoding/json"
	"fmt"

	file "github.com/up9inc/oas-diff/json"
)

type Info struct {
	Title          string   `json:"title" diff:"title"`
	Description    string   `json:"description,omitempty" diff:"description"`
	TermsOfService string   `json:"termsOfService,omitempty" diff:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" diff:"contact"`
	License        *License `json:"license,omitempty" diff:"license"`
	Version        string   `json:"version" diff:"version"`
}

type Contact struct {
	Name  string `json:"name,omitempty" diff:"name"`
	URL   string `json:"url,omitempty" diff:"url"`
	Email string `json:"email,omitempty" diff:"email"`
}

type License struct {
	Name string `json:"name" diff:"name"`
	URL  string `json:"url,omitempty" diff:"url"`
}

func ParseInfo(file file.JsonFile) (*Info, error) {
	var infoModel Info
	node := file.GetNodeData(OAS_INFO_KEY)
	if node != nil {
		err := json.Unmarshal(*node, &infoModel)
		if err != nil {
			return nil, fmt.Errorf("failed to Unmarshal Info struct: %v", err)
		}
	}

	return &infoModel, nil
}
