package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type stringsMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewStringsMapDiffer(opts DifferentiatorOptions) *stringsMapDiffer {
	return &stringsMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (s *stringsMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.StringsMap{}))
}

func (s *stringsMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.Loose {
		handleLooseMap[model.StringsMap](a, b)
	}

	return s.differ.DiffMap(path, a, b)
}

func (s *stringsMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
