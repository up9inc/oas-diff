package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type stringsMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewStringsMapDiffer(opts DifferentiatorOptions) *stringsMapDiffer {
	return &stringsMapDiffer{
		opts: opts,
	}
}

func (s *stringsMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.StringsMap{}))
}

func (s *stringsMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.Loose {
		handleLooseMap[model.StringsMap](a, b)
	}

	return df(path, a, b, parent)
}

func (s *stringsMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
