package validator

import (
	"encoding/json"
	"errors"

	"github.com/santhosh-tekuri/jsonschema/v5"
	file "github.com/up9inc/oas-diff/json"
)

const OAS_SCHEMA_URL = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"

type Validator interface {
	InitOAS31Schema() error
	Validate(jsonFile file.JsonFile) error
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

func (v *validator) InitOAS31Schema() error {
	jsonSchema := file.NewJsonFile("schema/OAS.json")
	err := jsonSchema.ValidatePath()
	if err != nil {
		return err
	}
	v.jsonSchema = jsonSchema

	sch, err := v.compiler.Compile(jsonSchema.GetPath())
	if err != nil {
		return err
	}
	v.schema = sch

	return nil
}

func (v *validator) Validate(jsonFile file.JsonFile) error {
	if v.schema == nil {
		err := v.InitOAS31Schema()
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
