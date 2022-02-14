package validator

import (
	file "github.com/up9inc/oas-diff/json"
)

const OAS_SCHEMA_URL = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"

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
