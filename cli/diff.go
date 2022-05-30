package cli

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	differentiator "github.com/up9inc/oas-diff/differentiator"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/reporter"
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
			TempDirFlag,
			BaseFileFlag,
			SecondFileFlag,
			TypeFilterFlag,
			HtmlOutputFlag,
			SummaryOutputFlag,
			LooseFlag,
			IncludeFilePathFlag,
			IgnoreDescriptionsFlag,
			IgnoreExamplesFlag,
		},
	}
}

func diffCmd(c *cli.Context) error {
	tempDir := c.String(TempDirFlag.Name)
	baseFilePath := c.String(BaseFileFlag.Name)
	secondFilePath := c.String(SecondFileFlag.Name)
	isHtmlOutput := c.Bool(HtmlOutputFlag.Name)
	isSummaryOutput := c.Bool(SummaryOutputFlag.Name)

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

	val := validator.NewValidator(tempDir)
	diff := differentiator.NewDifferentiator(val, differentiator.DifferentiatorOptions{
		TypeFilter:         strings.ToLower(c.String(TypeFilterFlag.Name)),
		Loose:              c.Bool(LooseFlag.Name),
		IncludeFilePath:    c.Bool(IncludeFilePathFlag.Name),
		IgnoreDescriptions: c.Bool(IgnoreDescriptionsFlag.Name),
		IgnoreExamples:     c.Bool(IgnoreExamplesFlag.Name),
	})

	changelog, err := diff.Diff(jsonFile, jsonFile2)
	if err != nil {
		return err
	}

	timeFormat := "15:04:05.000"
	if runtime.GOOS == "windows" {
		timeFormat = "15_04_05"
	}
	outputPath := fmt.Sprintf("%s_%s", "changelog", time.Now().Format(timeFormat))
	rep := reporter.NewJSONReporter(changelog)
	outputData, err := rep.Build()
	if err != nil {
		return err
	}

	jsonOutput := fmt.Sprintf("%s%s", outputPath, ".json")
	err = saveDiffOutputFile(jsonOutput, outputData)
	if err != nil {
		return err
	}

	if isSummaryOutput {
		rep = reporter.NewSummaryReporter(jsonFile, jsonFile2, changelog)
		outputData, err := rep.Build()
		if err != nil {
			return err
		}

		endpointsOutput := fmt.Sprintf("%s_%s%s", "summary", outputPath, ".json")
		err = saveDiffOutputFile(endpointsOutput, outputData)
		if err != nil {
			return err
		}
	}

	if isHtmlOutput {
		rep = reporter.NewHTMLReporter(changelog, "")
		outputData, err := rep.Build()
		if err != nil {
			return err
		}

		htmlOutput := fmt.Sprintf("%s%s", outputPath, ".html")
		err = saveDiffOutputFile(htmlOutput, outputData)
		if err != nil {
			return err
		}

		return openBrowser(htmlOutput)
	}

	return nil
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

func openBrowser(path string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", path).Start()
	case "windows":
		//err = exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
		err = exec.Command("cmd", "/C", "start", path).Start()
	case "darwin":
		err = exec.Command("open", path).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
