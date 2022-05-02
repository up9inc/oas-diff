package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type serverVariableDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewServerVariableDiffer(opts DifferentiatorOptions) *serverVariableDiffer {
	return &serverVariableDiffer{
		opts: opts,
	}
}

func (s *serverVariableDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ServerVariable{}))
}

func (s *serverVariableDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.ServerVariable)
		bValue, bOk := b.Interface().(model.ServerVariable)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (s *serverVariableDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
