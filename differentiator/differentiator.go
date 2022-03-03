package differentiator

import (
	"fmt"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (changeMap, error)
}

type DifferentiatorOptions struct {
	Loose               bool
	IncludeFilePath     bool
	ExcludeDescriptions bool
}

type differentiator struct {
	validator validator.Validator
	opts      DifferentiatorOptions
	differ    *lib.Differ

	info    *infoDiff
	servers *serversDiff
	paths   *pathsDiff
}

func NewDifferentiator(val validator.Validator, opts DifferentiatorOptions) Differentiator {
	// custon differs
	stringDiffer := NewStringDiffer(opts)
	parametersDiffer := NewParameterDiffer(opts)

	differ, err := lib.NewDiffer(lib.CustomValueDiffers(stringDiffer), lib.CustomValueDiffers(parametersDiffer), lib.StructMapKeySupport(), lib.DisableStructValues(), lib.SliceOrdering(false))
	if err != nil {
		panic(err)
	}

	stringDiffer.differ = differ
	parametersDiffer.differ = differ

	v := &differentiator{
		validator: val,
		opts:      opts,
		differ:    differ,
		info:      NewInfoDiff(),
		servers:   NewServersDiff(),
		paths:     NewPathsDiff(),
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
	err = d.info.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	changeMap[d.info.key] = d.info.changelog

	// servers
	err = d.servers.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	changeMap[d.servers.key] = d.servers.changelog

	// paths
	err = d.paths.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	changeMap[d.paths.key] = d.paths.changelog

	return changeMap, nil
}
