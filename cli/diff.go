package cli

import (
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
	infoModel, err := model.ParseInfo(jsonFile)
	if err != nil {
		return err
	}

	fmt.Println(infoModel)

	// info file 2
	infoModel2, err := model.ParseInfo(jsonFile2)
	if err != nil {
		return err
	}

	fmt.Println(infoModel2)

	// info diff
	infoChangelog, err := diff.Diff(infoModel, infoModel2)
	if err != nil {
		return err
	}

	buildOutput(model.OAS_INFO_KEY, infoChangelog, &sb)

	// servers file 1
	serversModel, err := model.ParseServers(jsonFile)
	if err != nil {
		return err
	}

	fmt.Println(serversModel)

	// servers file 2
	serversModel2, err := model.ParseServers(jsonFile2)
	if err != nil {
		return err
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
