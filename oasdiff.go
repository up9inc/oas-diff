package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
	cmd "github.com/up9inc/oas-diff/cli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "oas-diff",
		Version:  "0.1.0-alpha",
		Usage:    "Validate/Diff OAS 3.1 files",
		HideHelp: true,
		Commands: []*cli.Command{
			cmd.RegisterValidateCmd(),
			cmd.RegisterDiffCmd(),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(cmd.Red(fmt.Sprintf("%v", err)))
	}
}