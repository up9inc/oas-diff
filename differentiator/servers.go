package differentiator

import (
	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

type serversDiff struct {
	*internalDiff
	data  *model.Servers
	data2 *model.Servers
}

func (s *serversDiff) Diff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator) (*serversDiff, error) {
	var err error
	s = &serversDiff{
		internalDiff: NewInternalDiff(model.OAS_SERVERS_KEY),
	}

	// schema
	err = s.schema.Build(validator)
	if err != nil {
		return nil, err
	}

	// servers1
	s.data, err = model.ParseServers(jsonFile)
	if err != nil {
		return nil, err
	}

	// servers2
	s.data2, err = model.ParseServers(jsonFile2)
	if err != nil {
		return nil, err
	}

	// servers changelog
	s.changelog.Changelog, err = lib.Diff(s.data, s.data2)
	if err != nil {
		return nil, err
	}

	return s, nil
}
