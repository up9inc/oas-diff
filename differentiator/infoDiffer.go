package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*infoDiffer)(nil)

type infoDiffer struct {
	*internalDiff
	data  *model.Info
	data2 *model.Info

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewInfoDiffer() *infoDiffer {
	return &infoDiffer{
		internalDiff: NewInternalDiff(model.OAS_INFO_KEY),
		data:         &model.Info{},
		data2:        &model.Info{},
	}
}

func (i *infoDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
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

func (i *infoDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Info{}))
}

func (i *infoDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if i.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.Info)
		bValue, bOk := b.Interface().(model.Info)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (i *infoDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	i.DiffFunc = dfunc
}
