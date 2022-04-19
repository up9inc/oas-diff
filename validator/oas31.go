package validator

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
	file "github.com/up9inc/oas-diff/json"
)

const (
	OAS31_SCHEMA_URL  = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"
	OAS31_SCHEMA_FILE = "validator/oas31.json"
)

func (v *validator) InitSchemaFromFile(schemaFile file.JsonFile) error {
	err := schemaFile.ValidatePath()
	if err != nil {
		return err
	}

	v.jsonSchema = schemaFile
	v.schema, err = v.compiler.Compile(v.jsonSchema.GetPath())
	if err != nil {
		return err
	}

	return nil
}

func (v *validator) InitSchemaFromURL(url string) error {
	if len(url) == 0 {
		url = OAS31_SCHEMA_URL
	}

	var err error
	v.schema, err = v.compiler.Compile(url)
	if err != nil {
		return err
	}

	return nil
}

func (v *validator) GetSchemaProperty(key string) (*jsonschema.Schema, error) {
	if v.schema == nil {
		err := v.InitSchemaFromURL(OAS31_SCHEMA_URL)
		if err != nil {
			return nil, err
		}
	}

	p, ok := v.schema.Properties[key]
	if !ok {
		return nil, fmt.Errorf("failed to find schema property for key: %s", key)
	}

	if p.Ref != nil {
		return p.Ref, nil
	}

	return p, nil
}

func (v *validator) GetSchemaPropertyRequiredFields(key string) ([]string, error) {
	p, err := v.GetSchemaProperty(key)
	if err != nil {
		return nil, err
	}

	return p.Required, nil
}

func (v *validator) GetSchemaPropertyFields(key string) ([]string, error) {
	p, err := v.GetSchemaProperty(key)
	if err != nil {
		return nil, err
	}

	props := make([]string, 0)
	for k := range p.Properties {
		props = append(props, k)
	}

	return props, nil
}
