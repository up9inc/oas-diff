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

	info       *infoDiff
	servers    *serversDiffer
	paths      *pathsDiffer
	webhooks   *webhooksDiffer
	components *componentsDiffer
}

func NewDifferentiator(val validator.Validator, opts DifferentiatorOptions) Differentiator {
	// strings
	stringDiffer := NewStringDiffer(opts)

	// structs
	infoDiff := NewInfoDiff()
	componentsDiffer := NewComponentsDiffer()

	// slices
	serversDiffer := NewServersDiffer()
	parametersDiffer := NewParameterDiffer(opts)

	// maps
	schemasMapDiffer := NewSchemasMapDiffer(opts)
	pathsDiffer := NewPathsDiffer()
	webhooksDiffer := NewWebhooksDiffer()
	headersDiffer := NewHeadersDiffer(opts)
	responsesDiffer := NewResponsesMapDiffer(opts)
	contentMapDiffer := NewContentMapDiffer(opts)
	encodingMapDiffer := NewEncodingMapDiffer(opts)
	linksDiffer := NewLinksMapDiffer(opts)
	callbacksDiffer := NewCallbacksDiffer(opts)
	examplesDiffer := NewExamplesDiffer(opts)
	serverVariablesDiffer := NewServerVariablesDiffer(opts)

	differ, err := lib.NewDiffer(
		// strings
		lib.CustomValueDiffers(stringDiffer),
		// structs
		lib.CustomValueDiffers(componentsDiffer),
		// slices
		lib.CustomValueDiffers(serversDiffer),
		lib.CustomValueDiffers(parametersDiffer),
		// maps
		lib.CustomValueDiffers(schemasMapDiffer),
		lib.CustomValueDiffers(pathsDiffer),
		lib.CustomValueDiffers(webhooksDiffer),
		lib.CustomValueDiffers(headersDiffer),
		lib.CustomValueDiffers(responsesDiffer),
		lib.CustomValueDiffers(contentMapDiffer),
		lib.CustomValueDiffers(encodingMapDiffer),
		lib.CustomValueDiffers(linksDiffer),
		lib.CustomValueDiffers(callbacksDiffer),
		lib.CustomValueDiffers(examplesDiffer),
		lib.CustomValueDiffers(serverVariablesDiffer),
		// options
		lib.StructMapKeySupport(),
		lib.DisableStructValues(),
		lib.SliceOrdering(false))
	if err != nil {
		panic(err)
	}

	stringDiffer.differ = differ

	infoDiff.differ = differ
	componentsDiffer.differ = differ

	serversDiffer.differ = differ
	parametersDiffer.differ = differ

	schemasMapDiffer.differ = differ
	pathsDiffer.differ = differ
	webhooksDiffer.differ = differ
	headersDiffer.differ = differ
	responsesDiffer.differ = differ
	contentMapDiffer.differ = differ
	encodingMapDiffer.differ = differ
	linksDiffer.differ = differ
	callbacksDiffer.differ = differ
	examplesDiffer.differ = differ
	serverVariablesDiffer.differ = differ

	v := &differentiator{
		validator:  val,
		opts:       opts,
		differ:     differ,
		info:       infoDiff,
		servers:    serversDiffer,
		paths:      pathsDiffer,
		webhooks:   webhooksDiffer,
		components: componentsDiffer,
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

	// webhooks
	err = d.webhooks.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.webhooks.key] = d.webhooks.changelog

	// components
	err = d.components.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.components.key] = d.components.changelog

	return output, nil
}
