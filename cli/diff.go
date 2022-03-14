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
			BaseFileFlag,
			SecondFileFlag,
			HtmlOutputFlag,
			LooseFlag,
			IncludeFilePathFlag,
			ExcludeDescriptionsFlag,
		},
	}
}

func diffCmd(c *cli.Context) error {
	baseFilePath := c.String(BaseFileFlag.Name)
	secondFilePath := c.String(SecondFileFlag.Name)
	isHtmlOutput := c.Bool(HtmlOutputFlag.Name)

	jsonFile := file.NewJsonFile(baseFilePath)
	_, err := jsonFile.Read()
	if err != nil {
		return err
	}

	jsonFile2 := file.NewJsonFile(secondFilePath)
	_, err = jsonFile2.Read()
	if err != nil {
		return err
	}

	val := validator.NewValidator()
	diff := differentiator.NewDifferentiator(val, differentiator.DifferentiatorOptions{
		Loose:               c.Bool(LooseFlag.Name),
		IncludeFilePath:     c.Bool(IncludeFilePathFlag.Name),
		ExcludeDescriptions: c.Bool(ExcludeDescriptionsFlag.Name),
	})

	changelog, err := diff.Diff(jsonFile, jsonFile2)
	if err != nil {
		return err
	}

	var outputData []byte
	outputPath := fmt.Sprintf("%s_%s", "changelog", time.Now().Format("15:04:05.000"))
	if isHtmlOutput {
		// TODO: HTML Template logic
		outputPath = fmt.Sprintf("%s%s", outputPath, ".html")
	} else {
		outputPath = fmt.Sprintf("%s%s", outputPath, ".json")
		outputData, err = json.MarshalIndent(changelog, "", "\t")
		if err != nil {
			return err
		}
	}

	return saveDiffOutputFile(outputPath, outputData)
}

func saveDiffOutputFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	dirPath, err := filepath.Abs("./")
	if err != nil {
		return err
	}

	fmt.Println(Green(fmt.Sprintf("report saved: %s", fmt.Sprintf("%s/%s", dirPath, path))))

	return nil
}
