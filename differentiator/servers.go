package differentiator

import (
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*serversDiff)(nil)

type serversDiff struct {
	*internalDiff
	data  *model.Servers
	data2 *model.Servers
}

func NewServersDiff() *serversDiff {
	return &serversDiff{
		internalDiff: NewInternalDiff(model.OAS_SERVERS_KEY),
		data:         &model.Servers{},
		data2:        &model.Servers{},
	}
}

func (s *serversDiff) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts *DifferentiatorOptions) error {
	var err error

	// opts
	s.opts = opts

	// schema
	err = s.schema.Build(validator)
	if err != nil {
		return err
	}

	// servers1
	s.filePath = jsonFile.GetPath()
	err = s.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// servers2
	s.filePath2 = jsonFile2.GetPath()
	err = s.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// servers changelog
	changes, err := s.diff(s.data, s.data2)
	if err != nil {
		return err
	}

	// changelogs
	s.internalDiff.handleArrayChanges(s.data, s.data2, changes)

	return nil
}
