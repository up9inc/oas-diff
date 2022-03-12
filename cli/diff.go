package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

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
			LooseFlag,
			IncludeFilePathFlag,
			ExcludeDescriptionsFlag,
		},
	}
}

func diffCmd(c *cli.Context) error {
	isLoose := c.Bool(LooseFlag.Name)
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
	diff := differentiator.NewDifferentiator(val, differentiator.DifferentiatorOptions{
		Loose:               isLoose,
		IncludeFilePath:     c.Bool(IncludeFilePathFlag.Name),
		ExcludeDescriptions: c.Bool(ExcludeDescriptionsFlag.Name),
	})

	changelog, err := diff.Diff(jsonFile, jsonFile2)
	if err != nil {
		return err
	}

	output, err := json.MarshalIndent(changelog, "", "\t")
	if err != nil {
		panic(err)
	}

	outputPath := "changelog"
	if isLoose {
		outputPath = fmt.Sprintf("%s%s", outputPath, "-loose")
	}
	outputPath = fmt.Sprintf("%s_%s%s", outputPath, time.Now().Format("15:04:05.000"), ".json")

	err = ioutil.WriteFile(outputPath, output, 0644)
	if err != nil {
		return err
	}

	dirPath, err := filepath.Abs("./")
	if err != nil {
		return err
	}

	fmt.Println(Green(fmt.Sprintf("report saved: %s", fmt.Sprintf("%s/%s", dirPath, outputPath))))

	return nil
}
