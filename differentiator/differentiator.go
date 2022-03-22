package differentiator

import (
	"fmt"
	"time"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (*ChangelogOutput, error)
}

type DifferentiatorOptions struct {
	Loose               bool `json:"loose"`
	IncludeFilePath     bool `json:"include-file-path"`
	ExcludeDescriptions bool `json:"exclude-descriptions"`
}

type differentiator struct {
	validator validator.Validator
	opts      DifferentiatorOptions
	differ    *lib.Differ

	info    *infoDiff
	servers *serversDiffer
	paths   *pathsDiffer
}

func NewDifferentiator(val validator.Validator, opts DifferentiatorOptions) Differentiator {
	// custom differs
	stringDiffer := NewStringDiffer(opts)

	// slices
	serversDiffer := NewServersDiffer()
	parametersDiffer := NewParameterDiffer(opts)

	// maps
	pathsDiffer := NewPathsDiffer()
	headersDiffer := NewHeadersDiffer(opts)
	responsesDiffer := NewResponsesDiffer(opts)
	linksDiffer := NewLinksDiffer(opts)

	differ, err := lib.NewDiffer(
		lib.CustomValueDiffers(stringDiffer),
		lib.CustomValueDiffers(serversDiffer),
		lib.CustomValueDiffers(parametersDiffer),
		lib.CustomValueDiffers(pathsDiffer),
		lib.CustomValueDiffers(headersDiffer),
		lib.CustomValueDiffers(responsesDiffer),
		lib.CustomValueDiffers(linksDiffer),
		lib.StructMapKeySupport(),
		lib.DisableStructValues(),
		lib.SliceOrdering(false))
	if err != nil {
		panic(err)
	}

	stringDiffer.differ = differ
	serversDiffer.differ = differ
	parametersDiffer.differ = differ
	pathsDiffer.differ = differ
	headersDiffer.differ = differ
	responsesDiffer.differ = differ
	linksDiffer.differ = differ

	v := &differentiator{
		validator: val,
		opts:      opts,
		differ:    differ,
		info:      NewInfoDiff(),
		servers:   serversDiffer,
		paths:     pathsDiffer,
	}

	return v
}

func (d *differentiator) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (*ChangelogOutput, error) {
	start := time.Now()

	err := d.validator.Validate(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile.GetPath())
	}

	err = d.validator.Validate(jsonFile2)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file", jsonFile2.GetPath())
	}

	// output
	//changeMap := NewChangeMap(d.opts)
	output := NewChangelogOutput(start, jsonFile.GetPath(), jsonFile2.GetPath(), d.opts)

	// info
	err = d.info.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.info.key] = d.info.changelog

	// servers
	err = d.servers.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.servers.key] = d.servers.changelog

	// paths
	err = d.paths.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.paths.key] = d.paths.changelog

	return output, nil
}
