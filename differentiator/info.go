package differentiator

import (
	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type infoDiff struct {
	*internalDiff
	data  *model.Info
	data2 *model.Info
}

func (i *infoDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) (*infoDiff, error) {
	var err error
	i = &infoDiff{
		internalDiff: NewInternalDiff(model.OAS_INFO_KEY),
	}

	// schema
	err = i.schema.Build(validator)
	if err != nil {
		return nil, err
	}

	// info1
	i.data, err = model.ParseInfo(jsonFile)
	if err != nil {
		return nil, err
	}

	// info2
	i.data2, err = model.ParseInfo(jsonFile2)
	if err != nil {
		return nil, err
	}

	// info changelog
	i.changelog.Changelog, err = lib.Diff(i.data, i.data2)
	if err != nil {
		return nil, err
	}

	return i, nil
}
