package differentiator

import (
	"reflect"

	lib "github.com/r3labs/diff/v3"
	"github.com/up9inc/oas-diff/model"
)

type serverVariablesMapDiffer struct {
	opts DifferentiatorOptions

	DiffFunc (func(path []string, a, b reflect.Value, p interface{}) error)
}

func NewServerVariablesMapDiffer(opts DifferentiatorOptions) *serverVariablesMapDiffer {
	return &serverVariablesMapDiffer{
		opts: opts,
	}
}

func (s *serverVariablesMapDiffer) Match(a, b reflect.Value) bool {
	return lib.AreType(a, b, reflect.TypeOf(model.ServerVariablesMap{}))
}

func (s *serverVariablesMapDiffer) Diff(dt lib.DiffType, df lib.DiffFunc, cl *lib.Changelog, path []string, a, b reflect.Value, parent interface{}) error {
	if s.opts.IgnoreDescriptions {
		ignoreDescriptionsFromMaps[model.ServerVariablesMap](a, b)
	}

	if s.opts.Loose {
		handleLooseMap[model.ServerVariablesMap](a, b)
	}

	return df(path, a, b, parent)
}

func (s *serverVariablesMapDiffer) InsertParentDiffer(dfunc func(path []string, a, b reflect.Value, p interface{}) error) {
	s.DiffFunc = dfunc
}
