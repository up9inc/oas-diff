package differentiator

import (
	"reflect"
	"strings"

	lib "github.com/r3labs/diff/v2"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*serversDiffer)(nil)

type serversDiffer struct {
	*internalDiff
	data  *model.Servers
	data2 *model.Servers

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewServersDiffer() *serversDiffer {
	return &serversDiffer{
		internalDiff: NewInternalDiff(model.OAS_SERVERS_KEY),
		data:         &model.Servers{},
		data2:        &model.Servers{},
	}
}

func (s *serversDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
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
	changes, err := s.differ.Diff(s.data, s.data2)
	if err != nil {
		return err
	}

	// changelogs
	s.internalDiff.handleArrayChanges(s.data, s.data2, changes)

	return nil
}

func (s *serversDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Servers{}))
}

func (s *serversDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.Loose {
		if a.Kind() == reflect.Invalid {
			cl.Add(lib.CREATE, path, nil, lib.ExportInterface(b))
			return nil
		}

		if b.Kind() == reflect.Invalid {
			cl.Add(lib.DELETE, path, lib.ExportInterface(a), nil)
			return nil
		}

		if a.Kind() != b.Kind() {
			return lib.ErrTypeMismatch
		}

		aValue, aOk := a.Interface().(model.Servers)
		bValue, bOk := b.Interface().(model.Servers)

		if aOk && bOk {
			aIds := aValue.FilterIdentifiers()
			bIds := bValue.FilterIdentifiers()

			for _, a := range aIds {
				for _, b := range bIds {
					if a != nil && b != nil {
						if a.Name != b.Name && strings.EqualFold(a.Name, b.Name) {
							// we don't want this case sensitive identifier comparison
							// set lower case for both identifiers and keep comparing
							aValue[a.Index].URL = strings.ToLower(aValue[a.Index].URL)
							bValue[b.Index].URL = strings.ToLower(bValue[b.Index].URL)
						}
					}
				}
			}
		}
	}

	return s.differ.DiffSlice(path, a, b)
}

func (s *serversDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
