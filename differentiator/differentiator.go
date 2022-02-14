package differentiator

import (
	"fmt"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (*changelog, error)
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

func (d *differentiator) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (*changelog, error) {
	err := d.validator.InitOAS31Schema()
	if err != nil {
		return nil, err
	}

	err = d.validator.Validate(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile.GetPath())
	}

	err = d.validator.Validate(jsonFile2)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile2.GetPath())
	}

	// info
	err = d.infoDiff(jsonFile, jsonFile2)
	if err != nil {
		return nil, err
	}

	// servers
	d.servers, err = model.ParseServers(jsonFile)
	if err != nil {
		return nil, err
	}

	// servers2
	d.servers2, err = model.ParseServers(jsonFile2)
	if err != nil {
		return nil, err
	}

	// servers changelog
	d.serversChangelog, err = lib.Diff(d.servers, d.servers2)
	if err != nil {
		return nil, err
	}

	return d.info.changelog, nil
}
