package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type serverDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewServerDiffer(opts DifferentiatorOptions) *serverDiffer {
	return &serverDiffer{
		opts: opts,
	}
}

func (s *serverDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.Server{}))
}

func (s *serverDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreDescriptions {
		aValue, aOk := a.Interface().(model.Server)
		bValue, bOk := b.Interface().(model.Server)

		if aOk {
			aValue.IgnoreDescriptions()
		}

		if bOk {
			bValue.IgnoreDescriptions()
		}
	}

	return df(path, a, b, parent)
}

func (s *serverDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
