package differentiator

import (
	lib "github.com/r3labs/diff/v2"
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
		data:         &model.Info{},
		data2:        &model.Info{},
	}
}

func (i *infoDiff) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	i.opts = opts

	// differ
	i.differ = differ

	// schema
	err = i.schema.Build(validator)
	if err != nil {
		return err
	}

	// info1
	i.filePath = jsonFile.GetPath()
	err = i.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// info2
	i.filePath2 = jsonFile2.GetPath()
	err = i.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// info changelog
	changes, err := i.differ.Diff(i.data, i.data2)
	if err != nil {
		return err
	}

	// changelogs
	i.internalDiff.handleChanges(changes)

	return nil
}
