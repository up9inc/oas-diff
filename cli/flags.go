package cli

import "github.com/urfave/cli/v2"

var (
	FileFlag = &cli.StringFlag{
		Name:     "file",
		Usage:    "Path of the OAS 3.1 file",
		Required: true,
	}
	FileFlag2 = &cli.StringFlag{
		Name:     "file2",
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
)
