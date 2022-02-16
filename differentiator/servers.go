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
		data:         nil,
		data2:        nil,
	}
}

func (s *serversDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) (*internalDiff, error) {
	var err error

	// schema
	err = s.schema.Build(validator)
	if err != nil {
		return nil, err
	}

	// servers1
	s.filePath = jsonFile.GetPath()
	s.data, err = model.ParseServers(jsonFile)
	if err != nil {
		return nil, err
	}

	// servers2
	s.filePath2 = jsonFile2.GetPath()
	s.data2, err = model.ParseServers(jsonFile2)
	if err != nil {
		return nil, err
	}

	// servers changelog
	changes, err := s.diff(s.data, s.data2)
	if err != nil {
		return nil, err
	}

	// changelogs
	err = s.handleArrayChanges(s.data, s.data2, changes)
	if err != nil {
		return nil, err
	}

	return s.internalDiff, nil
}
