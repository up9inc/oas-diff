package cli

import "github.com/urfave/cli/v2"

var (
	BaseFileFlag = &cli.StringFlag{
		Name:     "base-file",
		Aliases:  []string{"f1"},
		Usage:    "path of the base OAS 3.1 file",
		Required: true,
	}
	SecondFileFlag = &cli.StringFlag{
		Name:     "second-file",
		Aliases:  []string{"f2"},
		Usage:    "path of the second OAS 3.1 file",
		Required: true,
	}
	TypeFilterFlag = &cli.StringFlag{
		Name:     "type",
		Aliases:  []string{"t"},
		Usage:    "Changelog Type filter (create/update/delete)",
		Required: false,
	}
	LooseFlag = &cli.BoolFlag{
		Name:     "loose",
		Usage:    "loosely diff, ignores global case sensitivity for strings comparisons and ignore headers that start with 'x-' and 'user-agent'",
		Required: false,
		Value:    false,
	}
	IncludeFilePathFlag = &cli.BoolFlag{
		Name:     "include-file-path",
		Usage:    "whether or not to include the full file path from the diff changelog",
		Required: false,
		Value:    false,
	}
	IgnoreDescriptionsFlag = &cli.BoolFlag{
		Name:     "ignore-descriptions",
		Usage:    "whether or not to ignore descriptions when performing the diff",
		Required: false,
		Value:    false,
	}
	IgnoreExamplesFlag = &cli.BoolFlag{
		Name:     "ignore-examples",
		Usage:    "whether or not to ignore examples when performing the diff",
		Required: false,
		Value:    false,
	}
	HtmlOutputFlag = &cli.BoolFlag{
		Name:     "html",
		Usage:    "save an html report",
		Required: false,
	}
)
