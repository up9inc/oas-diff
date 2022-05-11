package differentiator

import (
	"fmt"
	"time"

	lib "github.com/r3labs/diff/v3"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/validator"
)

type Differentiator interface {
	Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (*ChangelogOutput, error)
}

type DifferentiatorOptions struct {
	TypeFilter         string `json:"type-filter"`
	Loose              bool   `json:"loose"`
	IncludeFilePath    bool   `json:"include-file-path"`
	IgnoreDescriptions bool   `json:"ignore-descriptions"`
	IgnoreExamples     bool   `json:"ignore-examples"`
}

type differentiator struct {
	validator validator.Validator
	opts      DifferentiatorOptions
	differ    *lib.Differ

	info         *infoDiffer
	servers      *serversDiffer
	paths        *pathsMapDiffer
	webhooks     *webhooksMapDiffer
	components   *componentsDiffer
	security     *securityRequirementsDiffer
	tags         *tagsDiffer
	externalDocs *externalDocsDiffer
}

func NewDifferentiator(val validator.Validator, opts DifferentiatorOptions) Differentiator {
	if len(opts.TypeFilter) > 0 {
		if opts.TypeFilter != "create" && opts.TypeFilter != "update" && opts.TypeFilter != "delete" {
			panic(fmt.Errorf(`invalid type filter value "%s", must be %s/%s/%s`, opts.TypeFilter, lib.CREATE, lib.UPDATE, lib.DELETE))
		}
	}

	// strings
	stringDiffer := NewStringDiffer(opts)

	// structs
	infoDiffer := NewInfoDiffer()
	serverDiffer := NewServerDiffer(opts)
	serverVariableDiffer := NewServerVariableDiffer(opts)
	pathItemDiffer := NewPathItemDiffer(opts)
	componentsDiffer := NewComponentsDiffer()
	externalDocsDiffer := NewExternalDocsDiffer()
	operationDiffer := NewOperationDiffer(opts)
	parameterDiffer := NewParameterDiffer(opts)
	schemaDiffer := NewSchemaDiffer(opts)
	referenceDiffer := NewReferenceDiffer(opts)
	requestBodyDiffer := NewRequestBodyDiffer(opts)
	responseDiffer := NewResponseDiffer(opts)
	linkDiffer := NewLinkDiffer(opts)
	securitySchemeDiffer := NewSecuritySchemeDiffer(opts)
	tagDiffer := NewTagDiffer(opts)

	// slices
	serversDiffer := NewServersDiffer()
	parametersDiffer := NewParametersDiffer(opts)
	tagsDiffer := NewTagsDiffer()
	securityRequirementsDiffer := NewSecurityRequirementsDiffer()
	schemasSlicesDiffer := NewSchemasSliceDiffer(opts)

	// maps
	anyMapDiffer := NewAnyMapDiffer(opts)
	stringsMapDiffer := NewStringsMapDiffer(opts)
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
	securitySchemesMapDiffer := NewSecuritySchemesMapDiffer(opts)

	differ, err := lib.NewDiffer(
		// strings
		lib.CustomValueDiffers(stringDiffer),
		// structs
		lib.CustomValueDiffers(infoDiffer),
		lib.CustomValueDiffers(serverDiffer),
		lib.CustomValueDiffers(serverVariableDiffer),
		lib.CustomValueDiffers(pathItemDiffer),
		lib.CustomValueDiffers(componentsDiffer),
		lib.CustomValueDiffers(operationDiffer),
		lib.CustomValueDiffers(parameterDiffer),
		lib.CustomValueDiffers(schemaDiffer),
		lib.CustomValueDiffers(externalDocsDiffer),
		lib.CustomValueDiffers(referenceDiffer),
		lib.CustomValueDiffers(requestBodyDiffer),
		lib.CustomValueDiffers(responseDiffer),
		lib.CustomValueDiffers(linkDiffer),
		lib.CustomValueDiffers(securitySchemeDiffer),
		lib.CustomValueDiffers(tagDiffer),
		// slices
		lib.CustomValueDiffers(serversDiffer),
		lib.CustomValueDiffers(parametersDiffer),
		lib.CustomValueDiffers(tagsDiffer),
		lib.CustomValueDiffers(schemasSlicesDiffer),
		// maps
		lib.CustomValueDiffers(anyMapDiffer),
		lib.CustomValueDiffers(stringsMapDiffer),
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
		lib.CustomValueDiffers(securitySchemesMapDiffer),
		// options
		lib.StructMapKeySupport(),
		lib.DisableStructValues(),
		lib.SliceOrdering(false))

	if err != nil {
		panic(err)
	}

	v := &differentiator{
		validator:    val,
		opts:         opts,
		differ:       differ,
		info:         infoDiffer,
		servers:      serversDiffer,
		paths:        pathsMapDiffer,
		webhooks:     webhooksMapDiffer,
		components:   componentsDiffer,
		security:     securityRequirementsDiffer,
		tags:         tagsDiffer,
		externalDocs: externalDocsDiffer,
	}

	return v
}

func (d *differentiator) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile) (*ChangelogOutput, error) {
	start := time.Now()

	err := d.validator.Validate(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file: %s", jsonFile.GetPath(), err.Error())
	}

	err = d.validator.Validate(jsonFile2)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid 3.1 OAS file: %s", jsonFile2.GetPath(), err.Error())
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

	// security
	err = d.security.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.security.key] = d.security.changelog

	// tags
	err = d.tags.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.tags.key] = d.tags.changelog

	// externalDocs
	err = d.externalDocs.InternalDiff(jsonFile, jsonFile2, d.validator, d.opts, d.differ)
	if err != nil {
		return nil, err
	}
	output.Changelog[d.externalDocs.key] = d.externalDocs.changelog

	// filter by type
	if len(d.opts.TypeFilter) > 0 {
		output.Changelog = output.Changelog.FilterByType(d.opts.TypeFilter)
	}

	return output, nil
}
