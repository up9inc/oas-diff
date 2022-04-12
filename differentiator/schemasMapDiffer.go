package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type schemasMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSchemasMapDiffer(opts DifferentiatorOptions) *schemasMapDiffer {
	return &schemasMapDiffer{
		opts: opts,
	}
}

func (s *schemasMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.SchemasMap{}))
}

func (s *schemasMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreExamples {
		ignoreExamplesFromMaps[model.SchemasMap](a, b)
	}

	if s.opts.Loose {
		handleLooseMap[model.SchemasMap](a, b)
	}

	return df(path, a, b, parent)
}

func (s *schemasMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
