package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type schemasMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSchemasMapDiffer(opts DifferentiatorOptions) *schemasMapDiffer {
	return &schemasMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (s *schemasMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.SchemasMap{}))
}

func (s *schemasMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.Loose {
		handleLooseMap[model.SchemasMap](a, b)
	}

	return s.differ.DiffMap(path, a, b)
}

func (s *schemasMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
