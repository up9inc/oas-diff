package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	file "github.com/up9inc/oas-diff/json"
	"github.com/up9inc/oas-diff/model"
	"github.com/up9inc/oas-diff/validator"
)

// make sure we implement the InternalDiff interface
var _ InternalDiff = (*externalDocsDiffer)(nil)

type externalDocsDiffer struct {
	*internalDiff
	data  *model.ExternalDoc
	data2 *model.ExternalDoc

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewExternalDocsDiffer() *externalDocsDiffer {
	return &externalDocsDiffer{
		internalDiff: NewInternalDiff(model.OAS_EXTERNAL_DOCS_KEY),
		data:         &model.ExternalDoc{},
		data2:        &model.ExternalDoc{},
	}
}

func (e *externalDocsDiffer) InternalDiff(jsonFile file.JsonFile, jsonFile2 file.JsonFile, validator validator.Validator, opts DifferentiatorOptions, differ *lib.Differ) error {
	var err error

	// opts
	e.opts = opts

	// differ
	e.differ = differ

	// schema
	err = e.schema.Build(validator)
	if err != nil {
		return err
	}

	// externalDocs1
	e.filePath = jsonFile.GetPath()
	err = e.data.Parse(jsonFile)
	if err != nil {
		return err
	}

	// externalDocs2
	e.filePath2 = jsonFile2.GetPath()
	err = e.data2.Parse(jsonFile2)
	if err != nil {
		return err
	}

	// externalDocs changelog
	changes, err := e.differ.Diff(e.data, e.data2)
	if err != nil {
		return err
	}

	// changelogs
	e.internalDiff.handleChanges(changes)

	return nil
}

func (e *externalDocsDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ExternalDoc{}))
}

func (e *externalDocsDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if e.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.ExternalDoc)
		bValue, bOk := b.Interface().(model.ExternalDoc)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (e *externalDocsDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	e.DiffFunc = dfunc
}
