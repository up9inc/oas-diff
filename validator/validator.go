package validator

import (
	"encoding/json"
	"errors"

	"github.com/santhosh-tekuri/jsonschema/v5"
	file "github.com/up9inc/oas-diff/json"
)

type Validator interface {
	InitSchemaFromFile(schemaFile file.JsonFile) error
	InitSchemaFromURL(url string) error
	Validate(jsonFile file.JsonFile) error
	GetSchemaProperty(key string) (*jsonschema.Schema, error)
	GetSchemaPropertyRequiredFields(key string) ([]string, error)
	GetSchemaPropertyFields(key string) ([]string, error)
}

type validator struct {
	compiler   *jsonschema.Compiler
	schema     *jsonschema.Schema
	jsonSchema file.JsonFile
}

func NewValidator() Validator {
	v := &validator{
		compiler:   jsonschema.NewCompiler(),
		schema:     nil,
		jsonSchema: nil,
	}
	v.compiler.Draft = jsonschema.Draft2020

	return v
}

func (v *validator) Validate(jsonFile file.JsonFile) error {
	if v.schema == nil {
		err := v.InitSchemaFromFile(nil)
		if err != nil {
			return err
		}
	}

	if jsonFile == nil {
		return errors.New("json file is nil")
	}

	if jsonFile.GetData() == nil {
		return errors.New("json file data is nil")
	}

	var d interface{}
	if err := json.Unmarshal(*jsonFile.GetData(), &d); err != nil {
		return err
	}

	return v.schema.Validate(d)
}
