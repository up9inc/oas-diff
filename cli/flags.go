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
		Name:     "type-filter",
		Aliases:  []string{"tf"},
		Usage:    "changelog Type filter (create/update/delete)",
		Required: false,
	}
	LooseFlag = &cli.BoolFlag{
		Name:     "loose",
		Aliases:  []string{"l"},
		Usage:    "loosely diff, ignores global case sensitivity for strings comparisons and ignore headers that start with 'x-' and 'user-agent'",
		Required: false,
		Value:    false,
	}
	IncludeFilePathFlag = &cli.BoolFlag{
		Name:     "include-file-path",
		Aliases:  []string{"ifp"},
		Usage:    "whether or not to include the full file path from the diff changelog",
		Required: false,
		Value:    false,
	}
	IgnoreDescriptionsFlag = &cli.BoolFlag{
		Name:     "ignore-descriptions",
		Aliases:  []string{"id"},
		Usage:    "whether or not to ignore descriptions when performing the diff",
		Required: false,
		Value:    false,
	}
	IgnoreExamplesFlag = &cli.BoolFlag{
		Name:     "ignore-examples",
		Aliases:  []string{"ie"},
		Usage:    "whether or not to ignore examples when performing the diff",
		Required: false,
		Value:    false,
	}
	HtmlOutputFlag = &cli.BoolFlag{
		Name:     "output-html",
		Aliases:  []string{"oh"},
		Usage:    "save an html report",
		Required: false,
	}
	EndpointsOutputFlag = &cli.BoolFlag{
		Name:     "output-endpoint",
		Aliases:  []string{"oe"},
		Usage:    "endpoint based changelog output",
		Required: false,
	}
)
