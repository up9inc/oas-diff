package differentiator

import (
	"fmt"

	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (changeMap, error)
}

type differentiator struct {
	validator validator.Validator

	info    *infoDiff
	servers *serversDiff
}

func NewDiff(val validator.Validator) Differentiator {
	v := &differentiator{
		validator: val,
		info:      NewInfoDiff(),
		servers:   NewServersDiff(),
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
	changeMap[model.OAS_INFO_KEY] = d.info.changelog

	// servers
	err = d.servers.Diff(jsonFile, jsonFile2, d.validator)
	if err != nil {
		return nil, err
	}
	changeMap[model.OAS_SERVERS_KEY] = d.servers.changelog

	return changeMap, nil
}
