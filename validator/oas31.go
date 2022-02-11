package validator

import (
	"errors"

	"github.com/santhosh-tekuri/jsonschema/v5"
	file "github.com/up9inc/oas-diff/json"
)

const OAS_SCHEMA_URL = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"

type Validator interface {
	InitOAS31Schema() error
	Validate() error
}

type validator struct {
	json       *file.JsonFile
	compiler   *jsonschema.Compiler
	schema     *jsonschema.Schema
	jsonSchema *file.JsonFile
}

func NewValidator(json *file.JsonFile) Validator {
	v := &validator{
		json:       json,
		compiler:   jsonschema.NewCompiler(),
		schema:     nil,
		jsonSchema: nil,
	}
	v.compiler.Draft = jsonschema.Draft2020

	return v
}

func (v *validator) InitOAS31Schema() error {
	jsonSchema := file.NewJsonFile("../schema/OAS.json")
	err := jsonSchema.ValidatePath()
	if err != nil {
		return err
	}

	sch, err := v.compiler.Compile(jsonSchema.GetPath())
	if err != nil {
		return err
	}

	v.jsonSchema = &jsonSchema
	v.schema = sch

	return nil
}

func (v *validator) Validate() error {
	if v.schema == nil {
		err := v.InitOAS31Schema()
		if err != nil {
			return err
		}
	}

	if v.json == nil {
		return errors.New("json file is nil")
	}

	if v.json.GetData() == nil {
		return errors.New("json file data is nil")
	}

	/* 	var i interface{}
	   	if err := json.Unmarshal(jsonData, &v); err != nil {
	   		return nil, err
	   	}

	   	return jsonData, sch.Validate(v)
	*/
	return v.schema.Validate(nil)
}
