package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	differentiator "github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
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
			IncludeFilePath,
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

	val := validator.NewValidator()
	opts := &differentiator.DifferentiatorOptions{
		IncludeFilePath:     c.Bool(IncludeFilePath.Name),
		ExcludeDescriptions: c.Bool(ExcludeDescriptions.Name),
	}
	diff := differentiator.NewDiff(val, opts)

	changelog, err := diff.Diff(jsonFile, jsonFile2)
	if err != nil {
		return err
	}

	output, err := json.MarshalIndent(changelog, "", "\t")
	if err != nil {
		panic(err)
	}

	outputPath := "changelog.json"
	err = ioutil.WriteFile(outputPath, output, 0644)
	if err != nil {
		return err
	}

	fmt.Println(Green(fmt.Sprintf("report saved: %s", outputPath)))

	return nil
}
