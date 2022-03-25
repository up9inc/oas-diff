package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*securityRequirementsDiffer)(nil)

type securityRequirementsDiffer struct {
	*internalDiff
	data  *model.SecurityRequirements
	data2 *model.SecurityRequirements

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSecurityRequirementsDiffer() *securityRequirementsDiffer {
	return &securityRequirementsDiffer{
		internalDiff: NewInternalDiff(model.OAS_SECURITY_KEY),
		data:         &model.SecurityRequirements{},
		data2:        &model.SecurityRequirements{},
	}
}

func (s *securityRequirementsDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	s.opts = opts

	// differ
	s.differ = differ

	// schema
	err = s.schema.Build(validator)
	if err != nil {
		return err
	}

	// securityRequirements1
	s.filePath = jsonFile.GetPath()
	err = s.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// securityRequirements2
	s.filePath2 = jsonFile2.GetPath()
	err = s.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// securityRequirements1 changelog
	changes, err := s.differ.Diff(s.data, s.data2)
	if err != nil {
		return err
	}

	// changelogs
	s.internalDiff.handleChanges(changes)

	return nil
}
