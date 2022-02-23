package differentiator

import (
	"fmt"

	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (changeMap, error)
}

type DifferentiatorOptions struct {
	IncludeFilePath bool
}

type differentiator struct {
	validator validator.Validator
	opts      *DifferentiatorOptions

	info    *infoDiff
	servers *serversDiff
	paths   *pathsDiff
}

func NewDiff(val validator.Validator, opts *DifferentiatorOptions) Differentiator {
	v := &differentiator{
		validator: val,
		opts:      opts,
		info:      NewInfoDiff(),
		servers:   NewServersDiff(),
		paths:     NewPathsDiff(),
	}
	if v.opts == nil {
		v.opts = &DifferentiatorOptions{
			IncludeFilePath: false,
		}
	}

	return v
}

func (d *differentiator) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (changeMap, error) {
	err := d.validator.Validate(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile.GetPath())
	}

	err = d.validator.Validate(jsonFile2)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile2.GetPath())
	}

	// change map
	changeMap := NewChangeMap()

	// info
	err = d.info.Diff(jsonFile, jsonFile2, d.validator)
	if err != nil {
		return nil, err
	}
	changeMap[d.info.key] = d.info.changelog

	// servers
	err = d.servers.Diff(jsonFile, jsonFile2, d.validator)
	if err != nil {
		return nil, err
	}
	changeMap[d.servers.key] = d.servers.changelog

	// paths
	err = d.paths.Diff(jsonFile, jsonFile2, d.validator)
	if err != nil {
		return nil, err
	}
	changeMap[d.paths.key] = d.paths.changelog

	return changeMap, nil
}
