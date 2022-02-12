package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/up9inc/oas-diff/console"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/validator"
	"github.com/urfave/cli/v2"
)

func RegisterValidateCmd() *cli.Command {
	return &cli.Command{
		Name:    "validate",
		Aliases: []string{"v"},
		Usage:   "Validate file to OAS 3.1 schema",
		Action:  validateCmd,
		Flags:   []cli.Flag{FileFlag},
	}
}

func validateCmd(c *cli.Context) error {
	filePath := c.String(FileFlag.Name)

	jsonFile := file.NewJsonFile(filePath)
	_, err := jsonFile.Read()
	if err != nil {
		return err
	}

	validator := validator.NewValidator()
	err = validator.InitOAS31Schema()
	if err != nil {
		return err
	}

	err = validator.Validate(jsonFile)
	if err != nil {
		sb := strings.Builder{}

		var validationError *jsonschema.ValidationError
		if errors.As(err, &validationError) {
			output := validationError.BasicOutput()
			for _, e := range output.Errors {
				if len(e.InstanceLocation) > 0 {
					sb.WriteString(fmt.Sprintf("'%s' %s\n", e.InstanceLocation, e.Error))
				}
			}

		} else {
			sb.WriteString(fmt.Sprintf("%#v", err))
		}

		fmt.Println(console.Red(sb.String()))

		return nil
	}

	fmt.Println(console.Green("Valid OAS 3.1 file!"))

	return nil
}
