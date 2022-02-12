package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
	"github.com/urfave/cli/v2"
)

func RegisterDiffCmd() *cli.Command {
	return &cli.Command{
		Name:    "diff",
		Aliases: []string{"d"},
		Usage:   "Diff between two OAS 3.1 files",
		Action:  diffCmd,
		Flags: []cli.Flag{
			FileFlag,
			FileFlag2,
		},
	}
}

func buildOutput(key string, changelog diff.Changelog, sb *strings.Builder) {
	for _, c := range changelog {
		sb.WriteString(fmt.Sprintf("\nproperty: %s\npath: %s\ntype: %s\nfrom: %v\nto: %v\n", key, c.Path, c.Type, c.From, c.To))
	}
}

func diffCmd(c *cli.Context) error {
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
	infoData := jsonFile.GetNodeData(model.OAS_INFO_KEY)

	var infoModel model.Info

	if infoData != nil {
		fmt.Printf("info1: %s\n", string(*infoData))

		err = json.Unmarshal(*infoData, &infoModel)
		if err != nil {
			return err
		}
	}

	fmt.Println(infoModel)

	// info file 2
	infoData2 := jsonFile2.GetNodeData(model.OAS_INFO_KEY)

	var infoModel2 model.Info

	if infoData2 != nil {
		fmt.Printf("info2: %s\n", string(*infoData2))

		err = json.Unmarshal(*infoData2, &infoModel2)
		if err != nil {
			return err
		}
	}

	fmt.Println(infoModel2)

	// info diff
	infoChangelog, err := diff.Diff(infoModel, infoModel2)
	if err != nil {
		return err
	}

	buildOutput(model.OAS_INFO_KEY, infoChangelog, &sb)

	// servers file 1

	serversData := jsonFile.GetNodeData(model.OAS_SERVERS_KEY)

	var serversModel model.Servers

	if serversData != nil {
		fmt.Printf("servers1: %s\n", string(*serversData))

		err = json.Unmarshal(*serversData, &serversModel)
		if err != nil {
			return err
		}
	}

	fmt.Println(serversModel)

	// servers file 2
	serversData2 := jsonFile2.GetNodeData(model.OAS_SERVERS_KEY)

	var serversModel2 model.Servers

	if serversData2 != nil {
		fmt.Printf("servers2: %s\n", string(*serversData2))

		err = json.Unmarshal(*serversData2, &serversModel2)
		if err != nil {
			return err
		}
	}

	fmt.Println(serversModel2)

	// servers diff
	serversChangelog, err := diff.Diff(serversModel, serversModel2)
	if err != nil {
		return err
	}

	buildOutput(model.OAS_SERVERS_KEY, serversChangelog, &sb)

	fmt.Println(Green(sb.String()))

	return nil
}
