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
	FileFlag = &cli.StringFlag{
		Name:     "file",
		Usage:    "Path of the OAS 3.1 file",
		Required: true,
	}
	FileFlag2 = &cli.StringFlag{
		Name:     "file2",
		Usage:    "Path of the second OAS 3.1 file",
		Required: true,
	}
)

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

func validate(filePath string) (*[]byte, error) {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	//sch, err := compiler.Compile(OAS_SCHEMA_URL)
	sch, err := compiler.Compile("schema/OAS.json")
	if err != nil {
		return nil, err
	}

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &jsonData, err
	}

	var v interface{}
	if err := json.Unmarshal(jsonData, &v); err != nil {
		return &jsonData, err
	}

	return &jsonData, sch.Validate(v)
}

func validateCommand(c *cli.Context) error {
	filePath := c.String(FileFlag.Name)

	_, err := validate(filePath)
	if err != nil {
		sb := strings.Builder{}

		var validationError *jsonschema.ValidationError
		if errors.As(err, &validationError) {
			output := validationError.BasicOutput()
			for _, e := range output.Errors {
				if len(e.InstanceLocation) > 0 {
					sb.WriteString(console.Red(fmt.Sprintf("'%s' %s\n", e.InstanceLocation, e.Error)))
				}
			}

		} else {
			sb.WriteString(console.Red(fmt.Sprintf("%#v", err)))
		}

		fmt.Println(sb.String())

		return nil
	}

	fmt.Println(console.Green("Valid OAS 3.1 file!"))

	return nil
}

func diffCommand(c *cli.Context) error {
	filePath := c.String(FileFlag.Name)

	fileData, err := validate(filePath)
	if err != nil {
		return fmt.Errorf("%s is not a valid 3.1 OAS file", filePath)
	}

	filePath2 := c.String(FileFlag2.Name)
	fileData2, err := validate(filePath2)
	if err != nil {
		return fmt.Errorf("%s is not a valid 3.1 OAS file", filePath2)
	}

	fmt.Println(len(*fileData))
	fmt.Println(len(*fileData2))

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
				Flags:   []cli.Flag{FileFlag},
			},
			{
				Name:    "diff",
				Aliases: []string{"d"},
				Usage:   "Diff between two OAS 3.1 files",
				Action:  diffCommand,
				Flags: []cli.Flag{
					FileFlag,
					FileFlag2,
				},
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
