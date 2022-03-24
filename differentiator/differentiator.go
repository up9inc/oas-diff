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
	paths      *pathsMapDiffer
	webhooks   *webhooksMapDiffer
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
	pathsMapDiffer := NewPathsMapDiffer()
	webhooksMapDiffer := NewWebhooksMapDiffer()
	headersMapDiffer := NewHeadersMapDiffer(opts)
	parametersMapDiffer := NewParametersMapDiffer(opts)
	responsesMapDiffer := NewResponsesMapDiffer(opts)
	contentMapDiffer := NewContentMapDiffer(opts)
	encodingMapDiffer := NewEncodingMapDiffer(opts)
	linksMapDiffer := NewLinksMapDiffer(opts)
	callbacksMapDiffer := NewCallbacksMapDiffer(opts)
	examplesMapDiffer := NewExamplesMapDiffer(opts)
	serverVariablesMapDiffer := NewServerVariablesMapDiffer(opts)
	requestBodiesMapDiffer := NewRequestBodiesMapDiffer(opts)

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
		lib.CustomValueDiffers(pathsMapDiffer),
		lib.CustomValueDiffers(webhooksMapDiffer),
		lib.CustomValueDiffers(headersMapDiffer),
		lib.CustomValueDiffers(parametersMapDiffer),
		lib.CustomValueDiffers(responsesMapDiffer),
		lib.CustomValueDiffers(contentMapDiffer),
		lib.CustomValueDiffers(encodingMapDiffer),
		lib.CustomValueDiffers(linksMapDiffer),
		lib.CustomValueDiffers(callbacksMapDiffer),
		lib.CustomValueDiffers(examplesMapDiffer),
		lib.CustomValueDiffers(serverVariablesMapDiffer),
		lib.CustomValueDiffers(requestBodiesMapDiffer),
		// options
		lib.StructMapKeySupport(),
		lib.DisableStructValues(),
		lib.SliceOrdering(false))
	if err != nil {
		panic(err)
	}

	// strings
	stringDiffer.differ = differ
	// structs
	infoDiff.differ = differ
	componentsDiffer.differ = differ
	// slices
	serversDiffer.differ = differ
	parametersDiffer.differ = differ
	// maps
	schemasMapDiffer.differ = differ
	pathsMapDiffer.differ = differ
	webhooksMapDiffer.differ = differ
	headersMapDiffer.differ = differ
	parametersMapDiffer.differ = differ
	responsesMapDiffer.differ = differ
	contentMapDiffer.differ = differ
	encodingMapDiffer.differ = differ
	linksMapDiffer.differ = differ
	callbacksMapDiffer.differ = differ
	examplesMapDiffer.differ = differ
	serverVariablesMapDiffer.differ = differ
	requestBodiesMapDiffer.differ = differ

	v := &differentiator{
		validator:  val,
		opts:       opts,
		differ:     differ,
		info:       infoDiff,
		servers:    serversDiffer,
		paths:      pathsMapDiffer,
		webhooks:   webhooksMapDiffer,
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
