package differentiator

import (
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*infoDiff)(nil)

type infoDiff struct {
	*internalDiff
	data  *model.Info
	data2 *model.Info
}

func NewInfoDiff() *infoDiff {
	return &infoDiff{
		internalDiff: NewInternalDiff(model.OAS_INFO_KEY),
		data:         nil,
		data2:        nil,
	}
}

func (i *infoDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) (*internalDiff, error) {
	var err error

	// schema
	err = i.schema.Build(validator)
	if err != nil {
		return nil, err
	}

	// info1
	i.filePath = jsonFile.GetPath()
	i.data, err = model.ParseInfo(jsonFile)
	if err != nil {
		return nil, err
	}

	// info2
	i.filePath2 = jsonFile2.GetPath()
	i.data2, err = model.ParseInfo(jsonFile2)
	if err != nil {
		return nil, err
	}

	// info changelog
	changes, err := i.diff(i.data, i.data2)
	if err != nil {
		return nil, err
	}

	// changelogs
	err = i.handleChanges(changes)
	if err != nil {
		return nil, err
	}

	return i.internalDiff, nil
}
