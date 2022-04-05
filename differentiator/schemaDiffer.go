package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type schemaDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewSchemaDiffer(opts DifferentiatorOptions) *schemaDiffer {
	return &schemaDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (s *schemaDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Schema{}))
}

func (s *schemaDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreExamples {
		aValue, aOk := a.Interface().(model.Schema)
		bValue, bOk := b.Interface().(model.Schema)

		if aOk {
			aValue.IgnoreExamples()
		}

		if bOk {
			bValue.IgnoreExamples()
		}

	}

	return df(path, a, b, parent)
}

func (s *schemaDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
