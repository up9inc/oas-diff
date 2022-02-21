package differentiator

import (
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*pathsDiff)(nil)

type pathsDiff struct {
	*internalDiff
	data  model.Paths
	data2 model.Paths
}

func NewPathsDiff() *pathsDiff {
	return &pathsDiff{
		internalDiff: NewInternalDiff(model.OAS_PATHS_KEY),
		data:         model.Paths{},
		data2:        model.Paths{},
	}
}

func (p *pathsDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) error {
	var err error

	// schema
	err = p.schema.Build(validator)
	if err != nil {
		return err
	}

	// paths1
	p.filePath = jsonFile.GetPath()
	err = p.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// paths2
	p.filePath2 = jsonFile2.GetPath()
	err = p.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// paths changelog
	changes, err := p.diff(p.data, p.data2)
	if err != nil {
		return err
	}

	// changelogs
	return p.handleChanges(changes)
}
