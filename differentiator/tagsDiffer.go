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
var _ InternalDiff = (*tagsDiffer)(nil)

type tagsDiffer struct {
	*internalDiff
	data  *model.Tags
	data2 *model.Tags

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewTagsDiffer() *tagsDiffer {
	return &tagsDiffer{
		internalDiff: NewInternalDiff(model.OAS_TAGS_KEY),
		data:         &model.Tags{},
		data2:        &model.Tags{},
	}
}

func (t *tagsDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	t.opts = opts

	// differ
	t.differ = differ

	// schema
	err = t.schema.Build(validator)
	if err != nil {
		return err
	}

	// tags1
	t.filePath = jsonFile.GetPath()
	err = t.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// tags2
	t.filePath2 = jsonFile2.GetPath()
	err = t.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// tags changelog
	changes, err := t.differ.Diff(t.data, t.data2)
	if err != nil {
		return err
	}

	// changelogs
	t.internalDiff.handleArrayChanges(t.data, t.data2, changes)

	return nil
}

func (s *tagsDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Tags{}))
}

func (t *tagsDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if t.opts.Loose {
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

		aValue, aOk := a.Interface().(model.Tags)
		bValue, bOk := b.Interface().(model.Tags)

		if aOk && bOk {
			aIds := aValue.FilterIdentifiers()
			bIds := bValue.FilterIdentifiers()

			for _, a := range aIds {
				for _, b := range bIds {
					if a != nil && b != nil {
						if a.Name != b.Name && strings.EqualFold(a.Name, b.Name) {
							// we don't want this case sensitive identifier comparison
							// set lower case for both identifiers and keep comparing
							aValue[a.Index].Name = strings.ToLower(aValue[a.Index].Name)
							bValue[b.Index].Name = strings.ToLower(bValue[b.Index].Name)
						}
					}
				}
			}
		}
	}

	return t.differ.DiffSlice(path, a, b)
}

func (t *tagsDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	t.DiffFunc = dfunc
}
