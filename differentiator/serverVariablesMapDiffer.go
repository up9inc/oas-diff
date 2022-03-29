package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v2"
	"github.com/up9inc/oas-diff/model"
)

type serverVariablesMapDiffer struct {
	opts   DifferentiatorOptions
	differ *lib.Differ

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewServerVariablesMapDiffer(opts DifferentiatorOptions) *serverVariablesMapDiffer {
	return &serverVariablesMapDiffer{
		opts:   opts,
		differ: nil,
	}
}

func (s *serverVariablesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ServerVariablesMap{}))
}

func (s *serverVariablesMapDiffer) Diff(cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.Loose {
		handleLooseMap[model.ServerVariablesMap](a, b)
	}

	return s.differ.DiffMap(path, a, b)
}

func (s *serverVariablesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
