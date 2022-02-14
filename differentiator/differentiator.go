package differentiator

import (
	"fmt"

	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) ([]*changelog, error)
}

type differentiator struct {
	validator validator.Validator

	info    *infoDiff
	servers *serversDiff
}

func NewDiff() Differentiator {
	v := &differentiator{
		validator: validator.NewValidator(),
	}

	return v
}

func (d *differentiator) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) ([]*changelog, error) {
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

	// changelog
	changelog := make([]*changelog, 0)

	// info
	d.info, err = d.info.Diff(jsonFile, jsonFile2, d.validator)
	if err != nil {
		return nil, err
	}
	changelog = append(changelog, d.info.changelog)

	// servers
	d.servers, err = d.servers.Diff(jsonFile, jsonFile2, d.validator)
	if err != nil {
		return nil, err
	}
	changelog = append(changelog, d.servers.changelog)

	return changelog, nil
}
