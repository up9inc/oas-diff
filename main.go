package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gustavomassa/oas-diff/model"
	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/tidwall/gjson"
)

const (
	OAS_SCHEMA_URL   = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"
	OAS_INFO_KEY     = "info"
	OAS_SERVERS_KEY  = "servers"
	OAS_PATHS_KEY    = "paths"
	OAS_WEBHOOKS_KEY = "webhooks"
)

func isSliceOfUniqueItems(xs []interface{}) bool {
	s := len(xs)
	m := make(map[string]struct{}, s)
	for _, x := range xs {
		key, _ := json.Marshal(&x)
		m[string(key)] = struct{}{}
	}
	return s == len(m)
}

func getJsonPathData(jsonData []byte, path string) (result []byte, err error) {
	node := gjson.GetBytes(jsonData, path)
	if !node.Exists() {
		return result, fmt.Errorf("failed to find json path: %s", path)
	}
	if node.Index > 0 {
		result = jsonData[node.Index : node.Index+len(node.Raw)]
	} else {
		result = []byte(node.Raw)
	}
	return result, nil
}

func main() {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	//sch, err := compiler.Compile("schema/OAS.json")
	sch, err := compiler.Compile(OAS_SCHEMA_URL)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	jsonData, err := ioutil.ReadFile("test/shipping_valid.json")
	if err != nil {
		log.Fatal(err)
	}

	var v interface{}
	if err := json.Unmarshal(jsonData, &v); err != nil {
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

	infoSch := sch.Properties[OAS_INFO_KEY]
	serversSch := sch.Properties[OAS_SERVERS_KEY]
	pathsSch := sch.Properties[OAS_PATHS_KEY]
	webhooksSch := sch.Properties[OAS_WEBHOOKS_KEY]

	fmt.Println(infoSch.String())
	fmt.Println(serversSch.String())
	fmt.Println(pathsSch.String())
	fmt.Println(webhooksSch.String())

	if !gjson.ValidBytes(jsonData) {
		panic("invalid json")
	}

	infoData, err := getJsonPathData(jsonData, OAS_INFO_KEY)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(infoData))

	var infoModel model.Info
	err = json.Unmarshal(infoData, &infoModel)
	if err != nil {
		panic(err)
	}

	fmt.Println(infoModel)
}
