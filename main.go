package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/tidwall/gjson"
	"github.com/up9inc/oas-diff/console"
	"github.com/urfave/cli/v2"
)

const (
	OAS_SCHEMA_URL   = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/schemas/v3.1/schema.json"
	OAS_INFO_KEY     = "info"
	OAS_SERVERS_KEY  = "servers"
	OAS_PATHS_KEY    = "paths"
	OAS_WEBHOOKS_KEY = "webhooks"
)

var (
	PathFlag = &cli.StringFlag{
		Name:     "path",
		Aliases:  []string{"p"},
		Usage:    "Path of the OAS 3.1 file",
		Required: true,
	}
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

func validateCommand(c *cli.Context) error {
	path := c.String(PathFlag.Name)

	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	//sch, err := compiler.Compile(OAS_SCHEMA_URL)
	sch, err := compiler.Compile("schema/OAS.json")
	if err != nil {
		return err
	}

	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var v interface{}
	if err := json.Unmarshal(jsonData, &v); err != nil {
		return err
	}

	if err = sch.Validate(v); err != nil {
		sb := strings.Builder{}
		sb.WriteString("ERROR List: \n")

		var validationError *jsonschema.ValidationError
		if errors.As(err, &validationError) {
			output := validationError.BasicOutput()
			for _, e := range output.Errors {
				if len(e.InstanceLocation) > 0 {
					sb.WriteString(fmt.Sprintf("'%s' %s\n", e.InstanceLocation, e.Error))
				}
			}
		} else {
			sb.WriteString(fmt.Sprintf("%#v", err))
		}

		return errors.New(sb.String())
	}

	fmt.Println(console.Green("Valid OAS 3.1 file!"))

	return nil
}

func main() {
	app := &cli.App{
		Name:  "oas-diff",
		Usage: "Validate/Diff OAS 3.1 files",
		Commands: []*cli.Command{
			{
				Name:    "validate",
				Aliases: []string{"v"},
				Usage:   "Validate file to OAS 3.1 schema",
				Action:  validateCommand,
				Flags:   []cli.Flag{PathFlag},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(console.Red(fmt.Sprintf("%v", err)))
	}
}
