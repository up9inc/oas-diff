package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*componentsDiffer)(nil)

type componentsDiffer struct {
	*internalDiff
	data  *model.Components
	data2 *model.Components

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewComponentsDiffer() *componentsDiffer {
	return &componentsDiffer{
		internalDiff: NewInternalDiff(model.OAS_COMPONENTS_KEY),
		data:         &model.Components{},
		data2:        &model.Components{},
	}
}

func (c *componentsDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	c.opts = opts

	// differ
	c.differ = differ

	// schema
	err = c.schema.Build(validator)
	if err != nil {
		return err
	}

	// components1
	c.filePath = jsonFile.GetPath()
	err = c.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// components2
	c.filePath2 = jsonFile2.GetPath()
	err = c.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// components changelog
	changes, err := c.differ.Diff(c.data, c.data2)
	if err != nil {
		return err
	}

	// changelogs
	c.internalDiff.handleChanges(changes)

	return nil
}

func (c *componentsDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Components{}))
}

func (c *componentsDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if c.opts.IgnoreExamples {
		aValue, aOk := a.Interface().(model.Components)
		bValue, bOk := b.Interface().(model.Components)

		if aOk {
			aValue.IgnoreExamples()
		}

		if bOk {
			bValue.IgnoreExamples()
		}
	}

	return df(path, a, b, parent)
}

func (c *componentsDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	c.DiffFunc = dfunc
}
