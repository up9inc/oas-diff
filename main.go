package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

const OAS_SCHEMA_URL = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"

func main() {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	//sch, err := compiler.Compile("schema/OAS.json")
	sch, err := compiler.Compile(OAS_SCHEMA_URL)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile("test/shipping_invalid.json")
	if err != nil {
		log.Fatal(err)
	}

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatal(err)
	}

	if err = sch.Validate(v); err != nil {
		var validationError *jsonschema.ValidationError
		if errors.As(err, &validationError) {
			output := validationError.BasicOutput()
			/* 			b, _ := json.MarshalIndent(output, "", "  ")
			   			fmt.Println(string(b)) */

			for _, e := range output.Errors {
				//if len(e.InstanceLocation) > 0 && !strings.HasSuffix(e.Error, fmt.Sprintf("'/$defs%s'", e.InstanceLocation)) {
				if len(e.InstanceLocation) > 0 {
					fmt.Printf("ERROR: '%s' %s\n", e.InstanceLocation, e.Error)
				}
			}
		} else {
			log.Fatalf("%#v", err)
		}
	}
}
