package cli

import (
	differentiator "github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
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

func diffCmd(c *cli.Context) error {
	filePath := c.String(FileFlag.Name)
	filePath2 := c.String(FileFlag2.Name)

	jsonFile := file.NewJsonFile(filePath)
	_, err := jsonFile.Read()
	if err != nil {
		return err
	}

	jsonFile2 := file.NewJsonFile(filePath2)
	_, err = jsonFile2.Read()
	if err != nil {
		return err
	}

	diff := differentiator.NewDiff()

	return diff.Diff(jsonFile, jsonFile2)
}
