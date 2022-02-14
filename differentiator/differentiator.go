package differentiator

import (
	"fmt"
	"strings"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) error
}

type differentiator struct {
	validator validator.Validator

	info *infoDiff

	servers          *model.Servers
	servers2         *model.Servers
	serversChangelog lib.Changelog
}

func NewDiff() Differentiator {
	v := &differentiator{
		validator: validator.NewValidator(),
	}

	return v
}

func buildOutput(key string, changelog lib.Changelog, sb *strings.Builder) {
	for _, c := range changelog {
		sb.WriteString(fmt.Sprintf("\nproperty: %s\npath: %s\ntype: %s\nfrom: %v\nto: %v\n", key, c.Path, c.Type, c.From, c.To))
	}
}

func (d *differentiator) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) error {
	err := d.validator.InitOAS31Schema()
	if err != nil {
		return err
	}

	err = d.validator.Validate(jsonFile)
	if err != nil {
		return fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile.GetPath())
	}

	err = d.validator.Validate(jsonFile2)
	if err != nil {
		return fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile2.GetPath())
	}

	// info
	err = d.infoDiff(jsonFile, jsonFile2)
	if err != nil {
		return nil
	}

	fmt.Printf("info schema: %v\n", d.info.schema)

	sb := strings.Builder{}
	buildOutput(model.OAS_INFO_KEY, d.info.changelog.Changelog, &sb)

	// servers
	d.servers, err = model.ParseServers(jsonFile)
	if err != nil {
		return err
	}

	// servers2
	d.servers2, err = model.ParseServers(jsonFile2)
	if err != nil {
		return err
	}

	// servers changelog
	d.serversChangelog, err = lib.Diff(d.servers, d.servers2)
	if err != nil {
		return err
	}

	buildOutput(model.OAS_SERVERS_KEY, d.serversChangelog, &sb)

	// output
	fmt.Println(sb.String())

	return nil
}
