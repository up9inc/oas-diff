package cli

import "github.com/urfave/cli/v2"

var (
	BaseFileFlag = &cli.StringFlag{
		Name:     "base-file",
		Usage:    "Path of the base OAS 3.1 file",
		Required: true,
	}
	SecondFileFlag = &cli.StringFlag{
		Name:     "second-file",
		Usage:    "Path of the second OAS 3.1 file",
		Required: true,
	}
	LooseFlag = &cli.BoolFlag{
		Name:     "loose",
		Usage:    "loosely diff",
		Required: false,
		Value:    false,
	}
	IncludeFilePathFlag = &cli.BoolFlag{
		Name:     "include-file-path",
		Usage:    "Whether or not to include the full file path from the diff changelog",
		Required: false,
		Value:    false,
	}
	ExcludeDescriptionsFlag = &cli.BoolFlag{
		Name:     "exclude-descriptions",
		Usage:    "Whether or not to exclude descriptions from the diff changelog",
		Required: false,
		Value:    false,
	}
	HtmlOutputFlag = &cli.BoolFlag{
		Name:     "html",
		Usage:    "save the changelog file as a html report",
		Required: false,
	}
)
