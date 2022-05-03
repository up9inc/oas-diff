package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type schemasSliceDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSchemasSliceDiffer(opts DifferentiatorOptions) *schemasSliceDiffer {
	return &schemasSliceDiffer{
		opts: opts,
	}
}

func (s *schemasSliceDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.SchemasSlice{}))
}

func (s *schemasSliceDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreDescriptions {
		ignoreDescriptionsFromSlices[model.SchemasSlice](a, b)
	}

	if s.opts.IgnoreExamples {
		ignoreExamplesFromSlices[model.SchemasSlice](a, b)
	}

	return df(path, a, b, parent)
}

func (s *schemasSliceDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
