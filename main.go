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

	"github.com/r3labs/diff/v2"
	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	"github.com/up9inc/oas-diff/console"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
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

func validate(filePath string) ([]byte, error) {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	//sch, err := compiler.Compile(OAS_SCHEMA_URL)
	sch, err := compiler.Compile("schema/OAS.json")
	if err != nil {
		return nil, err
	}

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var v interface{}
	if err := json.Unmarshal(jsonData, &v); err != nil {
		return nil, err
	}

	return jsonData, sch.Validate(v)
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
					sb.WriteString(fmt.Sprintf("'%s' %s\n", e.InstanceLocation, e.Error))
				}
			}

		} else {
			sb.WriteString(fmt.Sprintf("%#v", err))
		}

		fmt.Println(console.Red(sb.String()))

		return nil
	}

	fmt.Println(console.Green("Valid OAS 3.1 file!"))

	return nil
}

func buildChangelog(key string, changelog diff.Changelog, sb *strings.Builder) {
	for _, c := range changelog {
		sb.WriteString(fmt.Sprintf("\nproperty: %s\npath: %s\ntype: %s\nfrom: %v\nto: %v\n", key, c.Path, c.Type, c.From, c.To))
	}
}

func diffCommand(c *cli.Context) error {
	filePath := c.String(FileFlag.Name)
	sb := strings.Builder{}

	//version := getJsonPathData(, OAS_INFO_KEY)

	jsonFile := file.NewJsonFile(filePath)
	_, err := jsonFile.Read()
	if err != nil {
		return err
	}

	validator := validator.NewValidator()
	err = validator.InitOAS31Schema()
	if err != nil {
		return err
	}

	err = validator.Validate(jsonFile)
	if err != nil {
		return fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile.GetPath())
	}

	filePath2 := c.String(FileFlag2.Name)
	jsonFile2 := file.NewJsonFile(filePath2)
	_, err = jsonFile2.Read()
	if err != nil {
		return err
	}

	err = validator.Validate(jsonFile2)
	if err != nil {
		return fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile2.GetPath())
	}

	// info file 1
	infoData := jsonFile.GetNodeData(OAS_INFO_KEY)

	var infoModel model.Info
	err = json.Unmarshal(*infoData, &infoModel)
	if err != nil {
		return err
	}

	fmt.Printf("info1: %s\n", string(*infoData))
	fmt.Println(infoModel)

	// info file 2
	infoData2 := jsonFile2.GetNodeData(OAS_INFO_KEY)

	var infoModel2 model.Info
	err = json.Unmarshal(*infoData2, &infoModel2)
	if err != nil {
		return err
	}

	fmt.Printf("info2: %s\n", string(*infoData2))
	fmt.Println(infoModel2)

	// info diff
	infoChangelog, err := diff.Diff(infoModel, infoModel2)
	if err != nil {
		return err
	}

	buildChangelog(OAS_INFO_KEY, infoChangelog, &sb)

	// servers file 1

	serversData := jsonFile.GetNodeData(OAS_SERVERS_KEY)

	var serversModel model.Servers
	err = json.Unmarshal(*serversData, &serversModel)
	if err != nil {
		return err
	}

	fmt.Printf("servers1: %s\n", string(*serversData))
	fmt.Println(serversModel)

	// servers file 2
	serversData2 := jsonFile2.GetNodeData(OAS_SERVERS_KEY)

	var serversModel2 model.Servers
	if serversData2 != nil {
		err = json.Unmarshal(*serversData2, &serversModel2)
		if err != nil {
			return err
		}
	}

	fmt.Printf("servers2: %s\n", string(*serversData2))
	fmt.Println(serversModel2)

	// servers diff
	serversChangelog, err := diff.Diff(serversModel, serversModel2)
	if err != nil {
		return err
	}

	buildChangelog(OAS_SERVERS_KEY, serversChangelog, &sb)

	fmt.Println(console.Green(sb.String()))

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
